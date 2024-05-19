*** Settings ***
Library  RequestsLibrary

*** Variables ***
${VALID_CLIENT_ID}    000000
${VALID_CLIENT_SECRET}    999999
${VALID_GRANT_TYPE}    client_credentials
${INVALID_CLIENT_ID}    000001
${INVALID_CLIENT_SECRET}    999998
${INVALID_GRANT_TYPE}    cliente_credencials
${INVALID_SCOPE}    resource.READY

*** Test Cases ***
Valid Client Credentials Request
  [Setup]    Setup Headers
  ${body}    Create Dictionary    grant_type=${VALID_GRANT_TYPE}    client_id=${VALID_CLIENT_ID}    client_secret=${VALID_CLIENT_SECRET}
  POST On Session    alias=session    url=/token  data=${body}    expected_status=anything
  Status Should Be    200

Invalid Client ID Request
  [Setup]    Setup Headers
  ${body}    Create Dictionary    grant_type=${VALID_GRANT_TYPE}    client_id=${INVALID_CLIENT_ID}    client_secret=${VALID_CLIENT_SECRET}
  POST On Session    alias=session    url=/token  data=${body}    expected_status=anything
  Status Should Be    401

Invalid Client Secret Request
  [Setup]    Setup Headers
  ${body}    Create Dictionary    grant_type=${VALID_GRANT_TYPE}    client_id=${VALID_CLIENT_ID}    client_secret=${INVALID_CLIENT_SECRET}
  POST On Session    alias=session    url=/token  data=${body}    expected_status=anything
  Status Should Be    401

Invalid Grant Type Request
  [Setup]    Setup Headers
  ${body}    Create Dictionary    grant_type=${INVALID_GRANT_TYPE}    client_id=${VALID_CLIENT_ID}    client_secret=${VALID_CLIENT_SECRET}
  POST On Session    alias=session    url=/token  data=${body}    expected_status=anything
  Status Should Be    400

Invalid Scope Request
  [Setup]    Setup Headers
  ${body}    Create Dictionary    scope=${INVALID_SCOPE}    grant_type=${VALID_GRANT_TYPE}    client_id=${VALID_CLIENT_ID}    client_secret=${VALID_CLIENT_SECRET}
  POST On Session    alias=session    url=/token  data=${body}    expected_status=anything
  Status Should Be    400

*** Keywords ***
Setup Headers
  ${content_type}    Create Dictionary    Content-Type=application/x-www-form-urlencoded
  Create Session    alias=session    url=http://localhost:9096    headers=${content_type}
