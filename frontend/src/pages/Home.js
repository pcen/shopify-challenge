import React, { useState } from 'react';

import { postJSON } from '../utils/requests';

// Home Page
const Home = props => {
  const [images, setImages] = useState({});

  const handleClick = () => {
    postJSON('/images', { image: 1 }).then(
      json => {
        setImages(json.images);
        console.log(json.images);
      },
      error => {
        setImages(error);
      }
    )
  }

  return (
    <React.Fragment>
      <h1>Home</h1>
      <button onClick={handleClick}>
        Get Image
      </button>
      <br /><br />
      <div>
        {Object.keys(images).length === 0 ?
          'No image data'
          :
          JSON.stringify(images)
        }
      </div>
    </React.Fragment>
  )
}

export default Home;
