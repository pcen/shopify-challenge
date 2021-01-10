import React, { useEffect, useState } from 'react';

import ImageLoader from '../components/ImageLoader';
import { postJSON } from '../utils/requests';

// Home Page
const Home = props => {
  const [images, setImages] = useState(null);

  const [query, setQuery] = useState('');
  const [includePublic, setIncludePublic] = useState(false);

  const submitQuery = () => {
    postJSON('/images', { query: query, includePublic: includePublic, }).then(
      json => {
        let loaders = [];
        json.images.forEach(image => {
          loaders.push(<ImageLoader id={image.ID} />)
        });
        setImages(loaders);
      },
      error => {
        console.log(error);
      }
    )
  }

  useEffect(() => {
    submitQuery();
  }, []);

  const handleClick = () => {
    submitQuery();
  }

  return (
    <React.Fragment>
      <h1>Home</h1>
      <button onClick={handleClick}>
        Search
      </button>
      <br /><br />
      <div>
        {images}
      </div>
    </React.Fragment>
  )
}

export default Home;
