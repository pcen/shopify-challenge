import React from 'react';
import { useState, useEffect } from 'react';

import Modal from './Modal';
import { getImage } from '../utils/requests';

import '../styles/gallery.css';

// Cache object URLs to image blobs instead of getting them from the backend
// on every new query. Cache is cleared upon loading the ImageGallery component
// to avoid stale cache entries.
const cache = new Map();

// EditImage
const EditImage = props => {
  const { metadata, submitChange } = props;
  const [editing, setEditing] = useState(false);
  const [data, setData] = useState(metadata);
  const [changes, setChanges] = useState({});
  const [changesMade, setChangesMade] = useState(false);

  // Set the initial changed metadata to be the origional metadata
  useEffect(() => {
    setChanges(metadata);
  }, []);

  // Update the image description
  const onChangeDescription = event => {
    setChanges({ ...changes, Description: event.target.value });
    setChangesMade(true);
  }

  // Update the image location
  const onChangeLocation = event => {
    setChanges({ ...changes, Geolocation: event.target.value });
    setChangesMade(true);
  }

  // Update the image visibility
  const onChangeVisibility = event => {
    setChanges({ ...changes, Private: event.target.checked });
    setChangesMade(true);
  }

  // On discard, set changes back to origional data
  const onDiscard = () => {
    setEditing(false);
    setChanges(data);
  }

  // On save, set origional to changed data
  const onSave = () => {
    setEditing(false);
    setData(changes);
  }

  const onClose = () => {
    if (changesMade) {
      submitChange(metadata.ID, changes)
    }
  }

  return (
    <Modal
      trigger={<div className='preview-button'>details</div>}
      onClose={onClose}
      content={
        <React.Fragment>
          <div className='edit-upload-header'>
            {editing ? `Editing ${metadata.Name}` : `${metadata.Name}`}
          </div>
          <div className='edit-upload-content'>
            <div className='edit-upload-input-region'>
              {/* Edit Visibility */}
              <div>
                {'Private '}
                <input type='checkbox'
                  defaultChecked={changes.Private}
                  onClick={editing ? onChangeVisibility : e => e.preventDefault()}
                />
              </div>
              <br></br>
              {/* Edit Description */}
              Description
              <textarea type='text'
                className='edit-upload-input'
                style={{ minHeight: '100px' }}
                value={changes.Description}
                onChange={editing ? onChangeDescription : () => { }}
              />
              <br></br>
              {/* Edit Location */}
              Location
              <input type='text'
                className='edit-upload-input'
                style={{ height: '20px' }}
                value={changes.Geolocation}
                onChange={editing ? onChangeLocation : () => { }}
              />
            </div>
            {/* Edit, Save, and Discard Changes */}
            <div className='edit-upload-actions'>
              <div
                className={!editing ? 'preview-button' : 'preview-button-disabled'}
                onClick={() => setEditing(true)}
              >
                Edit
              </div>
              <div className='preview-button' onClick={onDiscard}>
                Discard
              </div>
              <div className='preview-button' onClick={onSave}>
                Save
              </div>
            </div>
          </div>
        </React.Fragment>
      }
    />
  )
}

// ImageView
const ImageView = props => {
  const { image, onEdit, onDelete } = props;
  const [deleted, setDeleted] = useState(false);
  const [source, setSource] = useState(null);

  useEffect(() => {
    if (cache.has(image.ID)) {
      setSource(cache.get(image.ID));
    } else {
      getImage(image.ID).then(blob => {
        let image = blob === null ? null : URL.createObjectURL(blob, { autoRevoke: false });
        setSource(image);
        if (!cache.has(image.ID) && image !== null) {
          cache.set(image.ID, image);
        }
      });
    }
  }, []);

  // Callback function when image is deleted
  const deleteImage = () => {
    if (!deleted) {
      setDeleted(true);
      URL.revokeObjectURL(cache.get(image.ID))
      cache.delete(image.ID);
      onDelete(image.ID);
    }
  }

  if (source === null) {
    return null;
  } else {
    return (
      <div className='preview-frame'>
        <img className='image' src={source} alt={`image ${image.ID}`}></img>
        <div className='image-title'>{image.Name}</div>
        <div className='preview-buttons'>
          <EditImage metadata={image} submitChange={onEdit} />
          <div className='preview-button' onClick={deleteImage}>delete</div>
        </div>
      </div>
    )
  }
}

// ImageGallery provides a gallery view of images obtained from querying the
// backend. The image blobs are cached in a cache which is cleared in the
// constructor.
class ImageGallery extends React.Component {
  constructor(props) {
    super(props);
    // Clear the image cache when image gallery component is loaded. This
    // ensures that the cache is not persistant across different
    // ImageGallery instances.
    for (let url of cache.values()) {
      URL.revokeObjectURL(url);
    }
    cache.clear();
  }

  render() {
    return this.props.metadata === null ? null : (
      <div className='image-gallery'>
        {Array.from(this.props.metadata.values(), image => {
          return (
            <ImageView
              image={image}
              onEdit={this.props.onEdit}
              onDelete={this.props.onDelete}
              key={image.ID}
            />
          )
        })}
      </div>
    )
  }
}

export default ImageGallery;
