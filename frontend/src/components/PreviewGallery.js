import React, { useState, useEffect } from 'react';
import Popup from 'reactjs-popup';

// EditImage component is a button that will open a popup when clicked,
// providing text input fields to modify metadata about the image metadata
// prop. This component wraps a reatjs-popup Popup component:
// https://react-popup.elazizi.com/
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

  // Submit changes made to the image upon closing the popup
  const handleClose = () => {
    submitChange(changes);
  }

  return (
    <Popup
      trigger={<div className='preview-button'>metadata</div>}
      onClose={handleClose}
      modal
      position='top center'
      closeOnDocumentClick={false}
      closeOnEscape={false}
    >
      {close => (
        <div className='edit-details-modal'>
          <button className="close" onClick={close}>&times;</button>
          {/* Title of Image */}
          <div className="header">
            {editing ? `Editing '${metadata.name}'` : `${metadata.name}`}
          </div>
          <div className="content">
            <div className='edit-input-region'>
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
                className='edit-input'
                style={{ minHeight: '100px' }}
                value={changes.description}
                onChange={editing ? onChangeDescription : () => { }}
              />
              <br></br>
              {/* Edit Location */}
              Location
              <input type='text'
                className='edit-input'
                style={{ height: '20px' }}
                value={changes.location}
                onChange={editing ? onChangeLocation : () => { }}
              />
            </div>
            {/* Edit, Save, and Discard Changes */}
            <div className='edit-actions'>
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
        </div>
      )}
    </Popup>
  )
}

// ImagePreview component displays a preview of the image, as well as buttons
// to edit the image metadata or remove the image from the set of images it
// belongs to.
const ImagePreview = props => {
  const { image, onChange, onDelete } = props;

  const deleteImage = () => {
    onDelete(image.name);
  }

  return (
    <div className='preview-frame'>
      <img className='preview-img' src={image.source} alt={image.name} />
      <div className='preview-buttons'>
        <EditImage metadata={image} submitChange={onChange} />
        <div className='preview-button' onClick={deleteImage}>remove</div>
      </div>
    </div>
  )
}

// PreviewGallery component takes a Map images and returns a grid with a
// preview of each image.
const PreviewGallery = props => {
  const { images, removeImage, editImage } = props;

  if (images.length === 0) {
    return null;
  }

  let previews = [];
  for (let i of images.values()) {
    previews.push(
      <ImagePreview
        image={i}
        onDelete={removeImage}
        onChange={editImage}
        key={i.name}
      />
    );
  }

  return (
    <div className='image-preview-container'>
      {previews}
    </div>
  )
}

export default PreviewGallery;
