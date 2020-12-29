// Utility functions for sending HTTP requests

// checkAuthStatus checks to see if the backend response is a 401 unauthorized
// response. If a 401 is recieved and the cause is an invalid token, then the
// user data is removed from local storage and the page is reloaded to clear
// the Redux store. Since there is no longer a user JWT, reloading the page
// will also redirect to the login page.
const checkAuthStatus = response => {
  if (!response.ok && response.status === 401) {
    response.json().then(json => {
      if (json.error === 'token invalid') {
        localStorage.removeItem('user');
        window.location.reload();
      }
    })
  }
}

// Posts the 'payload' object to the given endpoint as JSON.
// jwt is an optional parameter that is set as the Authorization if it is
// passed to the function call.
async function postJSON(endpoint, payload, jwt) {
  let headers = {
    'Content-Type': 'application/json',
  };
  if (jwt !== undefined) {
    headers.Authorization = jwt;
  }
  const response = await fetch(endpoint, {
    method: 'POST',
    headers,
    body: JSON.stringify(payload),
  });
  checkAuthStatus(response);
  return response.json();
}

// Gets response as JSON from the given endpoint.
// jwt is an optional parameter that is set as the Authorization if it is
// passed to the function call.
async function get(endpoint, jwt) {
  let headers = {};
  if (jwt !== undefined) {
    headers.Authorization = jwt;
  }
  const response = await fetch(endpoint, {
    method: 'GET',
    headers,
  });
  checkAuthStatus(response);
  return response.json();
}

export { get, postJSON }
