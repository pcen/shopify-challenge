// Utility functions for sending HTTP requests

// checkAuthStatus checks to see if the backend response is a 401 unauthorized
// response. If a 401 is recieved, the user data is removed from local storage
// and the page is reloaded to clear the Redux store. Since there is no longer
// a user JWT, reloading the page will also redirect to the login page.
const checkAuthStatus = (response) => {
  if (!response.ok && response.status === 401) {
    localStorage.removeItem('user');
    window.location.reload();
    return false;
  }
  return true;
}

// handleJSON returns the JSON response from a request, or a Promise.reject on
// errors.
const handleJSON = response => {
  return response.json().then(json => {
    checkAuthStatus(response);
    // Returns backend error on 400 response
    if (response.ok && response.status === 400) {
      return Promise.reject(json.error);
    }
    // Returns backend error on 403 response
    if (!response.ok && response.status === 403) {
      return Promise.reject(json.error);
    }
    return json;
  }).catch(err => {
    if (!err) {
      return Promise.reject("response body invalid");
    } else {
      return Promise.reject(err);
    }
  })
}

// handleBlob returns the blob response from a request, or null on errors.
const handleBlob = response => {
  return response.blob().then(blob => {
    checkAuthStatus(response)
    // Returns null on errors
    if (!response.ok) {
      return null;
    }
    return blob;
  }).catch(() => {
    return null;
  })
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
  return handleJSON(response);
}

// getImage gets the image with the specified id from the backend. Returns
// image as a blob if valid, otherwise returns null.
async function getImage(id) {
  const headers = addAuth({
    'Cache-Control': 'no-cache',
  });
  let endpoint = '/image/'.concat(id.toString());
  const response = await fetch(endpoint, {
    method: 'GET',
    headers,
  });
  return handleBlob(response);
}

// Gets response as JSON from the given endpoint.
async function get(endpoint) {
  const headers = addAuth({});
  const response = await fetch(endpoint, {
    method: 'GET',
    headers,
  });
  return handleJSON(response);
}

// Creates multipart form data for a given image.
const imagesFormData = images => {
  let fd = new FormData();
  let meta = [];
  for (let image of images.values()) {
    fd.set(image.name, image.file, image.name);
    meta.push({
      name: image.name,
      description: image.description,
      location: image.location,
      private: image.private,
      format: image.file.type,
    });
  }
  fd.set('meta', JSON.stringify(meta));
  return fd;
}

// Posts images to the given endpoint as multipart form data. Image metadata is
// sent as a JSON string.
async function postImages(endpoint, images) {
  // Browser will add Content-Type header
  // https://muffinman.io/blog/uploading-files-using-fetch-multipart-form-data/
  const headers = addAuth({});
  const response = await fetch(endpoint, {
    method: 'POST',
    headers,
    body: imagesFormData(images),
  })
  return handleJSON(response);
}

export { get, postJSON, postImages, getImage }
