import React from 'react';

import '../styles/upload.css';

// UploadButton component is a button that allows for uploading of multiple
// files.
const UploadButton = props => {
  const { title, onUpload } = props;

  return (
    <React.Fragment>
      <input id='upload-btn' type='file' multiple hidden onChange={onUpload} />
      <label className='upload-single' htmlFor='upload-btn'>{title}</label>
    </React.Fragment>
  )
}

export default UploadButton;
