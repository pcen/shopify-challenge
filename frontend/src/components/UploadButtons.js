import React from 'react';

import '../styles/upload.css';

// UploadButton component is a button that allows for uploading of multiple
// files.
const UploadToClient = props => {
  const { title, onUpload } = props;
  return (
    <React.Fragment>
      <input id='upload-btn' type='file' multiple hidden onChange={onUpload} />
      <label className='upload-btn' htmlFor='upload-btn'>{title}</label>
    </React.Fragment>
  )
}

const UploadToBackend = props => {
  const { title, onUpload } = props;
  return (
    <label className='upload-btn' onClick={onUpload}>{title}</label>
  )
}

export { UploadToClient, UploadToBackend };
