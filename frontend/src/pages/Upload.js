import React, { useState } from 'react';

import UploadButton from '../components/UploadButton';
import ImagesPreview from '../components/ImagesPreview';

import '../styles/images.css';

// Upload Images Page
const Upload = props => {
  const [images, setImages] = useState(new Map());
  const [update, setUpdate] = useState(false);

  const handleUpload = event => {
    let added = false;
    if (event.target.files.length === 0) {
      return;
    }
    const files = event.target.files;
    for (let i = 0; i < files.length; i++) {
      const f = files.item(i);
      f.source = URL.createObjectURL(f);
      if (!images.has(f.name)) {
        images.set(f.name, f);
        added = true;
      }
    }
    if (added) {
      setUpdate(!update);
    }
  }

  return (
    <React.Fragment>
      <h1>Upload</h1>
      <div>Upload images to preview and add them to repository</div>
      <br />
      <UploadButton title='Choose Images' onUpload={handleUpload}></UploadButton>
      <br /><br />
      <ImagesPreview previewImages={images} />
    </React.Fragment>
  )
}

export default Upload;
