import { useState, useEffect } from 'react';

import { getImage } from '../utils/requests';

import '../styles/gallery.css';

const ImageLoader = props => {
  const [source, setSource] = useState(null);

  const { id } = props;

  useEffect(() => {
    if (id === null) {
      return;
    }
    getImage(id).then(blob => {
      let image = blob === null ? null : URL.createObjectURL(blob);
      setSource(image);
    });
  }, [id]);

  if (source === null || id === null) {
    return null;
  }

  return (
    <img className='image' src={source} alt={id}></img>
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
      <img className='preview-img' src={image.source} alt={image.name} />
      <div className='preview-buttons'>
        {/* <EditImage metadata={image} submitChange={onChange} /> */}
        <div className='preview-button' onClick={deleteImage}>delete</div>
      </div>
    </div>
  )
}

const ImageGallery = props => {
  const { metadata } = props;

  // handle changes made to an existing image
  const handleChange = (id, changes) => {
    console.log('changing image', id);
  }

  // handle deleting an image
  const handleDelete = id => {
    console.log('deleting image', id);
  }

  return metadata === null ? null : (
    <div className='image-gallery'>
      {Array.from(metadata, image => {
        return (
          <ImageView
            image={image}
            onChange={handleChange}
            onDelete={handleDelete}
            key={image.ID}
          />
        )
      })}
    </div>
  )
}

export default ImageGallery;
