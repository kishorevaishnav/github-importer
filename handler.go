package main

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strings"
)

//go:embed static/*.html
var content embed.FS

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFS(content, "static/*.html"))

	renderHTML(w, tmpl, "index.html", nil)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the form data
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit for file size
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	accessToken := r.FormValue("accessToken")
	userName := r.FormValue("userName")
	repoName := r.FormValue("repoName")

	// Get the uploaded file
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Unable to get uploaded file. Err: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Read the content of the file into a bytes.Buffer
	var fileBuffer bytes.Buffer
	_, err = io.Copy(&fileBuffer, file)
	if err != nil {
		http.Error(w, "Unable to read file", http.StatusInternalServerError)
		return
	}

	// Actual business starts here...

	// Read the XLSX & retrieve all the values.
	testCases := readXLSX(fileBuffer.Bytes())

	var ghURLs []string
	for _, tc := range testCases {
		if tc.Title == "Title" {
			continue
		}
		ghDescription := fmt.Sprintf("**Description:**\n%s\n\n**Test Steps:**\n%s\n\n**Test Data:**\n%s\n\n**Expected Results:**\n%s\n", tc.Description, tc.TestSteps, tc.TestData, tc.ExpectedResults)
		fmt.Printf("tcTitle -%s, ghDescription -%s, tcLabels -%s, tcAssignee -%s", tc.Title, ghDescription, tc.Labels, tc.Assignee)
		ghInit(accessToken, userName, repoName)
		ghURL, err := ghSubmitIssue(tc.Title, ghDescription, tc.Labels, tc.Assignee)
		if err != nil {
			http.Error(w, "Error creating issue: "+err.Error(), http.StatusInternalServerError)
			return
		}
		ghURLs = append(ghURLs, ghURL)
	}

	// Success everything went well.
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Success: Your request was processed successfully.\n\nHere are the Github Issue URLs:\n"+strings.Join(ghURLs, "\n"))
}

func renderHTML(w http.ResponseWriter, t *template.Template, tmplFile string, data interface{}) {
	err := t.ExecuteTemplate(w, tmplFile, data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
