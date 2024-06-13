# token generator

### 1. Create Credentials

- Login into [cloud](https://console.cloud.google.com) console and create a project if don't have one.
- Go to `APIs & Services` and enable the Gmail, Calendar, Google Docs and Drive APIs,
- Go to `OAuth consent screen` and setup the concent screen.
- Go to `Credentials > Create Credentials > OAuth client ID` and create credentials by selecting `Web Application` type. Make sure `http://localhost:7000/callback-gl` is at the top in the list of Authorized redirect URIs.
  <img width="1754" alt="Screenshot 2022-12-21 at 7 50 14 PM" src="https://user-images.githubusercontent.com/86593105/208926980-ee571d75-d250-46bb-a3ec-196984292d5e.png">

### 2. Generate Tokens

- Download the newly created credentials as JSON.
- Clone the [token generator](https://github.com/RiteshHarihar/token-generator) repo and put the credential file at the root of the project. Rename the credential file to `credentials.json`.
- Run the project and follow the steps. Tokens generated to a new file called `token.json`.
