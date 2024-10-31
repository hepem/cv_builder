# CV Builder

Allows to build a CV based on a JSON data.

### Requirements
- Go

### How to use
First remove the `.template` of the file `data/cv_data.json.template` and complete it.

Then Run:
```bash
go run main.go cv_structs.go
```

### Future improvements:
- Write output pdfs on english & spanish
- Different templates to be used
- Anonymize the pdf data