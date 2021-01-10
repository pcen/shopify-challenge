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

export default ImageLoader;
