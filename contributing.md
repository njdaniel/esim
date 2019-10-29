# Contributing to Esim
These are mostly guidelines to contritubing to this project.
Use your best judgement and feel free to propose changes to this document in a 
pull request.

#### Table of Contents


[Where to ask Questions](#i-have-a-question)

[What to know before I get started](#what-to-know-before-i-get-started)

[How can I contribute?](#how-can-i-contribute)
 * [Reporting bugs](#reporting-bugs)
 * [Suggesting Enhancements](#suggesting-enhancements)
 * [Your First Code Contribution](#your-first-code-contribution)
 * [Pull Requests](#pull-requests)
 
[Styleguides](#styleguides)


## Where to ask Questions

Join slack: 
https://join.slack.com/t/rtp-gophers/shared_invite/enQtNTE3NjIyMTgyODgyLTRhOTcxODBlNjc3NGYxNTI4Mzg5Mzg4OTY5ZjVmOGQ4ZDIyZjIxMzUxZDJlNjc0MzNmZjM2MmI3YmRlMzFjOTk
Ask question in the #go-projects channels.

## What to know before I get started

### Dependency management

Project uses [Go modules](https://golang.org/cmd/go/#hdr-Modules__module_versions__and_more) to manage dependencies on external packages. This requires a working Go environment with version 1.13 or greater installed.

All dependencies are vendored in the `vendor/` directory.

To add or update a new dependency, use the `go get` command:

```bash
# Pick the latest tagged release.
go get example.com/some/module/pkg

# Pick a specific version.
go get example.com/some/module/pkg@vX.Y.Z
```

Tidy up the `go.mod` and `go.sum` files and copy the new/updated dependency to the `vendor/` directory:


```bash
# The GO111MODULE variable can be omitted when the code isn't located in GOPATH.
GO111MODULE=on go mod tidy

GO111MODULE=on go mod vendor
```

You have to commit the changes to `go.mod`, `go.sum` and the `vendor/` directory before submitting the pull request.
 
 
## How can I contribute?

### Steps to Contribute
Should you wish to work on an issue, please claim it first by commenting on the GitHub issue that you want to work on it. 
This is to prevent duplicated efforts from contributors on the same issue.

Fork project and create new branch.
Create and reference a 'Draft Pull Request' in the Issue and mark with WIP until ready to be merged.  

For quick compile and testing
```
#running the project
go run main.go

#run tests
make test
```

### Reporting Bugs
Follow this guideline to help contributors to 
understand your report, reproduce the behavior, and find related reports.

Before creating a bug report check under Issues to see if bug is already 
created.

When creating a bug report please include as many details as possible.
Fill out the [required template](https://github.com/atom/.github/blob/master/.github/ISSUE_TEMPLATE/bug_report.md).

### Suggesting Enhancements

### Your First Code Contribution

### Pull Request Checklist




 
## Styleguides

