import React, { useEffect, useState } from 'react';

import ImageGallery from '../components/ImageGallery';
import SearchBar from '../components/SearchBar';
import { postJSON } from '../utils/requests';

// Home Page
const Home = props => {
  const [images, setImages] = useState(null);

  const [includePublic, setIncludePublic] = useState(false);

  const submitQuery = queryString => {
    postJSON('/images', { query: queryString, includePublic: includePublic, }).then(
      json => {
        setImages(json.images);
      },
      error => {
        console.log(error);
      }
    )
  }

  const handleQuery = queryString => {
    submitQuery(queryString);
  }

  useEffect(() => {
    submitQuery('');
  }, []);

  return (
    <React.Fragment>
      <h1>Home</h1>
      <div
        style={{
          display: 'flex',
          flexDirection: 'row',
          alignItems: 'center',
          justifyContent: 'center',
        }}
      >
        <SearchBar onSubmit={handleQuery} />
        <input
          style={{ margin: '0vw 5px 0vw 3vw' }}
          type='checkbox'
          defaultChecked={includePublic}
          onClick={() => { setIncludePublic(!includePublic) }}
        />
        <div>Include Public Images</div>
      </div>

      <br /><br />
      <div>
        <ImageGallery metadata={images} />
      </div>
    </React.Fragment >
  )
}

export default Home;
