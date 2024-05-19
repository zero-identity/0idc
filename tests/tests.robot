*** Settings ***
Library  RequestsLibrary

*** Variables ***
${CLIENT_ID}    000000
${CLIENT_SECRET}    999999
${BODY}    grant_type=client_credentials&client_id=${CLIENT_ID}&client_secret=${CLIENT_SECRET}

*** Test Cases ***
Client Credentials Test
  # ${client_auth}    Create List    000000    999999
  ${headers}    Create Dictionary    Content-Type=application/x-www-form-urlencoded
  Create Session    alias=session    url=http://localhost:9096    headers=${headers}
  POST On Session    alias=session    url=/token  data=${BODY}
