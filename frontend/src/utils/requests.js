// Utility functions for sending HTTP requests

// checkAuthStatus checks to see if the backend response is a 401 unauthorized
// response. If a 401 is recieved and the cause is an invalid token, then the
// user data is removed from local storage and the page is reloaded to clear
// the Redux store. Since there is no longer a user JWT, reloading the page
// will also redirect to the login page.
const checkAuthStatus = response => {
  return response.json().then(json => {
    if (!response.ok && response.status === 401) {
      if (json && json.error === 'token invalid') {
        localStorage.removeItem('user');
        window.location.reload();
      }
    }
    return json;
  }).catch(() => {
    return Promise.reject("response body is not JSON");
  });
}

// addAuth adds an 'Authorization' header to the passed headers object if a
// user object exists in local storage.
const addAuth = headers => {
  const user = JSON.parse(localStorage.getItem('user'));
  if (user != null) {
    return {
      ...headers,
      Authorization: user.authToken,
    };
  } else {
    return headers;
  }
}

// Posts the 'payload' object to the given endpoint as JSON.
async function postJSON(endpoint, payload) {
  const headers = addAuth({
    'Content-Type': 'application/json',
  });
  const response = await fetch(endpoint, {
    method: 'POST',
    headers,
    body: JSON.stringify(payload),
  });
  return checkAuthStatus(response);
}

// Gets response as JSON from the given endpoint.
async function get(endpoint) {
  const headers = addAuth({});
  const response = await fetch(endpoint, {
    method: 'GET',
    headers,
  });
  return checkAuthStatus(response);
}

// Creates multipart form data for a given image.
const imageFormData = image => {
  let fd = new FormData();
  fd.set('image', image.file, image.name);
  fd.set('description', image.description);
  fd.set('location', image.location);
  fd.set('private', image.private);
  return fd;
}

// Posts an image to the given endpoint as multipart form data.
async function postImage(endpoint, image) {
  // Browser will add Content-Type header
  // https://muffinman.io/blog/uploading-files-using-fetch-multipart-form-data/
  const headers = addAuth({});
  const response = await fetch(endpoint, {
    method: 'POST',
    headers,
    body: imageFormData(image),
  })
  return checkAuthStatus(response);
}

export { get, postJSON, postImage }
