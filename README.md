# google-analytics-v4-example

[![License](http://img.shields.io/:license-mit-blue.svg?style=flat-square)](LICENSE)

## Overview

This is an example of a simple console application that pulls pageview data from Google Analytics.

## Pre-requisites

### Enabling Google Analytics Reporting API

To query Google Analytics data, you'll need to enable the [Google Analytics Reporting API](https://console.cloud.google.com/apis/library/analyticsreporting.googleapis.com) in your Google Cloud Platform project.

### Creating a service account

1. Create a [service account](https://console.cloud.google.com/iam-admin/serviceaccounts) in Google Cloud Platform console for your project.
1. Assign the service account no permissions/roles.
1. Click "Create Key" to create a private key and save it in JSON format.
1. In Google Analytics, open Admin > View > View User Management and add the email address of the service account you just created (it will have an email like `[name]@[project ID].iam.gserviceaccount.com`.
1. Grant the user only "Read & Analyze" permissions.

### Finding your Google Analytics View ID

1. In Google Analytics, open Admin > View > View Settings
1. Save the View ID displayed

## Usage

```bash
go run main.go -- --viewID 123456 --keyFile google-analytics-viewer-key.json
```
