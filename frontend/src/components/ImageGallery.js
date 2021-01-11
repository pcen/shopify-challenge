import { useState, useEffect } from 'react';

import { getImage } from '../utils/requests';

import '../styles/gallery.css';

// Cache object URLs to image blobs instead of getting them from the backend
// on every new query. Cache is cleared upon loading the ImageGallery component
// to avoid stale cache entries.
const cache = new Map();

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

  useEffect(() => {
    // Clear the image cache when image gallery component is loaded. This
    // ensures that the cache is not persistant across different
    // ImageGallery instances.
    for (let url of cache.values()) {
      URL.revokeObjectURL(url);
    }
    cache.clear();
  }, [])

  // handle changes made to an existing image
  const handleChange = (id, changes) => {
    console.log('changing image', id);
  }

  // handle deleting an image
  const handleDelete = id => {
    console.log('deleting image', id);
    console.log(cache);
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
