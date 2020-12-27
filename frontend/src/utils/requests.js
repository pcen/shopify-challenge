// Utility functions for sending HTTP requests

// Posts the 'payload' object to the given endpoint as JSON
async function postJSON(endpoint, payload) {
  const response = await fetch(endpoint, {
    method: 'POST',
    mode: 'cors',
    cache: 'no-cache',
    credentials: 'same-origin',
    headers: {
      'Content-Type': 'application/json',
    },
    redirect: 'follow',
    referrerPolicy: 'no-referrer',
    body: JSON.stringify(payload),
  });
  return response.json();
}

export { postJSON }
