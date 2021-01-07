import React, { useState } from 'react';

import ImageLoader from '../components/ImageLoader';

// Home Page
const Home = props => {
  const [images, setImages] = useState({});

  const [imageID, setImageID] = useState(null);

  const handleClick = () => {
    setImageID(1);
    console.log('set image id');
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
        <ImageLoader id={imageID} />
      </div>

    </React.Fragment>
  )
}

export default Home;
