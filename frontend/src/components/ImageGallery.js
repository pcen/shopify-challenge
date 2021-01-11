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

  // Set the initial changed metadata to be the origional metadata
  useEffect(() => { setChanges(metadata); }, []);

  // Update the image description
  const onChangeDescription = event => {
    setChanges({ ...changes, description: event.target.value });
    console.log(event.target.value);
  }

  // Update the image location
  const onChangeLocation = event => {
    setChanges({ ...changes, location: event.target.value });
  }

  // Update the image visibility
  const onChangeVisibility = event => {
    setChanges({ ...changes, private: event.target.checked });
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

  console.log(metadata);

  return (
    <Modal
      trigger={<div className='preview-button'>details</div>}
      onClose={() => { submitChange(changes) }}
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
                  defaultChecked={changes.private}
                  onClick={editing ? onChangeVisibility : e => e.preventDefault()}
                />
              </div>
              <br></br>
              {/* Edit Description */}
              Description
              <textarea type='text'
                className='edit-upload-input'
                style={{ minHeight: '100px' }}
                value={changes.description}
                onChange={editing ? onChangeDescription : () => { }}
              />
              <br></br>
              {/* Edit Location */}
              Location
              <input type='text'
                className='edit-upload-input'
                style={{ height: '20px' }}
                value={changes.location}
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

// ImageLoader
const ImageLoader = props => {
  const [source, setSource] = useState(null);

  const { id } = props;

  useEffect(() => {
    if (cache.has(id)) {
      setSource(cache.get(id));
    } else {
      getImage(id).then(blob => {
        let image = blob === null ? null : URL.createObjectURL(blob, { autoRevoke: false });
        setSource(image);
        if (!cache.has(id) && image !== null) {
          cache.set(id, image);
        }
      });
    }
  }, []);

  return (
    <React.Fragment>
      {
        source === null ? null :
          <img className='image' src={source} alt={id}></img>
      }
    </React.Fragment>
  );
}

// ImageView
const ImageView = props => {
  const { image, onChange, onDelete } = props;

  const deleteImage = () => {
    onDelete(image.ID);
  }

  return (
    <div className='preview-frame'>
      <ImageLoader id={image.ID} />
      <div className='preview-buttons'>
        <EditImage metadata={image} submitChange={onChange} />
        <div className='preview-button' onClick={deleteImage}>delete</div>
      </div>
    </div>
  )
}

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

  // handle changes made to an existing image
  handleChange = (id, changes) => {
    console.log('changing image', id);
  }

  // handle deleting an image
  handleDelete = id => {
    console.log('deleting image', id);
    console.log(cache);
  }

  render() {
    return this.props.metadata === null ? null : (
      <div className='image-gallery'>
        {Array.from(this.props.metadata, image => {
          return (
            <ImageView
              image={image}
              onChange={this.handleChange}
              onDelete={this.handleDelete}
              key={image.ID}
            />
          )
        })}
      </div>
    )
  }
}

export default ImageGallery;
