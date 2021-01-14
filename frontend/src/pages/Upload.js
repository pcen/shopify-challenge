import React, { useState, useReducer } from 'react';
import { useHistory } from 'react-router-dom';

import { UploadToBackend, UploadToClient } from '../components/UploadButtons';
import PreviewGallery from '../components/PreviewGallery';
import { postImages } from '../utils/requests';

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
    file: file,
  }
}

// Upload Images Page
const Upload = props => {
  const [images, setImages] = useState(new Map());
  const [ignored, forceUpdate] = useReducer(x => x + 1, 0);
  const [hasUploaded, setHasUploaded] = useState(false);

  const history = useHistory();

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
      forceUpdate();
    }
  }

  const handleRemove = name => {
    if (images.delete(name)) {
      forceUpdate();
    }
  }

  const handleEdit = changes => {
    images.set(changes.name, changes);
    forceUpdate();
  }

  const handleSend = () => {
    if (images.size === 0 || hasUploaded) {
      return;
    }
    setHasUploaded(true);
    postImages('/upload', images).then(
      json => {
        console.log(json);
        history.push('/home');
      },
      error => {
        console.log(error);
        setHasUploaded(false);
      }
    );
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
      <UploadToClient
        title='Choose Images'
        onUpload={handleUpload}
      />
      <UploadToBackend
        title='Upload'
        onUpload={handleSend}
      />
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
