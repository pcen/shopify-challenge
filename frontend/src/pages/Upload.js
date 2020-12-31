import React, { useState } from 'react';

import UploadButton from '../components/UploadButton';
import PreviewGallery from '../components/PreviewGallery';

import '../styles/images.css';

// newImageMetadata returns an image metadata object given a file object.
const newImageMetadata = file => {
  return {
    name: file.name,
    description: '',
    location: '',
    private: true,
    source: URL.createObjectURL(file),
    type: file.type,
  }
}

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
      if (!images.has(f.name)) {
        images.set(f.name, newImageMetadata(f));
        added = true;
      }
    }
    if (added) {
      setUpdate(!update);
    }
  }

  const handleRemove = name => {
    if (images.delete(name)) {
      setUpdate(!update);
    }
  }

  const handleEdit = changes => {
    images.set(changes.name, changes);
    setUpdate(!update);
  }

  return (
    <React.Fragment>
      <h1 style={{ margin: '5px 0px 5px 0px' }}>Upload Images</h1>
      <div style={{ padding: '0px 40px 0px 40px' }}>
        Upload images to preview and add them to your repository. To add
        information about images, click metadata and then edit to make
        changes.
      </div>
      <br></br>
      <UploadButton title='Choose Images' onUpload={handleUpload}></UploadButton>
      <br /><br />
      <PreviewGallery
        images={images}
        removeImage={handleRemove}
        editImage={handleEdit}
      />
    </React.Fragment>
  )
}

export default Upload;
