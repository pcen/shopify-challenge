import React, { useState, useEffect, useReducer } from 'react';
import { useStore } from 'react-redux';

import Modal from './Modal';
import { getReq, getImage } from '../utils/requests';

import '../styles/gallery.css';

// Cache object URLs to image blobs instead of getting them from the backend
// on every new query. Cache is cleared upon loading the ImageGallery component
// to avoid stale cache entries.
const cache = new Map();

// TagList component lists all of the tags associated with an image.
const TagList = props => {
  const { tags } = props;

  return (
    <React.Fragment>
      <div style={{ fontSize: '13pt' }}>Tags</div>
      <br></br>
      <div className='image-tags'>
        {Array.from(tags.split(','), (v, i) => {
          return <div className='image-tag' key={i}>{v}</div>
        })}
      </div>
    </React.Fragment>
  )
}

// EditImageModalButtons component is the set of buttons used to edit, save
// changes, and discard changes for an image.
const EditImageModalButtons = props => {
  const { editing, setEditing, onDiscard, onSave } = props;
  return (
    <div className='edit-upload-actions'>
      <div
        className={!editing ? 'preview-button' : 'preview-button-selected'}
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
  )
}

// EditImage
const EditImage = props => {
  const { metadata, submitChange } = props;
  const [editing, setEditing] = useState(false);
  const [data, setData] = useState(metadata);
  const [changes, setChanges] = useState({});
  const [changesMade, setChangesMade] = useState(false);
  const [ignore, forceUpdate] = useReducer(x => x + 1, 0);
  const user = useStore().getState().user;

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
    setChangesMade(false);
  }

  // On save, set origional to changed data
  const onSave = () => {
    setEditing(false);
    setData(changes);
  }

  const handleClose = () => {
    if (changesMade) {
      submitChange(metadata.ID, changes)
    }
    setEditing(false);
  }

  // On open, check if the image has been tagged since query result metadata
  // was received from the backend.
  const checkForTags = () => {
    console.log(metadata.MLTags);
    if (metadata.MLTags === '') {
      getReq(`image/${metadata.ID}/tags`).then(
        json => {
          metadata.MLTags = json.tags;
          forceUpdate();
        },
        error => {
          console.log(error);
        }
      )
    }
  }

  return (
    <Modal
      trigger={<div className='preview-button' onClick={checkForTags}>details</div>}
      onOpen={checkForTags}
      onClose={handleClose}
      content={
        <React.Fragment>
          <div className='edit-upload-header'>
            {editing ? `Editing ${metadata.Name}` : `${metadata.Name}`}
          </div>
          <div className='edit-image-content'>
            <div className='image-content-left'>
              <div className='edit-image-input-image'>
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
                  className='edit-image-input'
                  style={{ minHeight: '100px' }}
                  value={changes.Description}
                  onChange={editing ? onChangeDescription : () => { }}
                />
                <br></br>
                {/* Edit Location */}
                Location
                <input type='text'
                  className='edit-image-input'
                  style={{ height: '20px' }}
                  value={changes.Geolocation}
                  onChange={editing ? onChangeLocation : () => { }}
                />
              </div>
              {/* Edit, Save, and Discard Changes */}
              {user.id !== metadata.UserID ? null :
                <EditImageModalButtons
                  editing={editing}
                  setEditing={setEditing}
                  onDiscard={onDiscard}
                  onSave={onSave}
                />
              }
            </div>
            <div className='image-content-right'>
              <TagList tags={metadata.MLTags} />
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
  const user = useStore().getState().user;

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
  }, [image.ID]);

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
        <img className='image' src={source} alt={`${image.ID}`}></img>
        <div className='image-title'>{image.Name}</div>
        <div className='preview-buttons'>
          <EditImage metadata={image} submitChange={onEdit} />
          {user.id !== image.UserID ?
            <div className='non-button'>public image</div>
            :
            <div className='preview-button' onClick={deleteImage}>delete</div>
          }
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
