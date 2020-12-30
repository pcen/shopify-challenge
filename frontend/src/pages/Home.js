import React, { useState } from 'react';

import { get } from '../utils/requests';

// Home Page
const Home = props => {
  const [images, setImages] = useState({});

  const handleClick = () => {
    get('/images').then(json => {
      setImages(json.images);
      console.log(json.images);
    })
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
          images['image 1']
        }
      </div>
    </React.Fragment>
  )
}

export default Home;
