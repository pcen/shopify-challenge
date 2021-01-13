import React, { useEffect, useReducer, useState } from 'react';

import ImageGallery from '../components/ImageGallery';
import SearchBar from '../components/SearchBar';
import { deleteReq, postJSON } from '../utils/requests';

// Home Page
const Home = props => {
  const [images, setImages] = useState(null);
  const [ignored, forceUpdate] = useReducer(x => x + 1, 0);

  const [includePublic, setIncludePublic] = useState(false);

  const submitQuery = queryString => {
    postJSON('/images', { query: queryString, includePublic: includePublic, }).then(
      json => {
        let result = new Map();
        for (let image of json.images) {
          result.set(image.ID, image);
        }
        setImages(result);
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

  const handleEdit = (id, change) => {
    images.set(id, change);
    postJSON(`/image/${id}/edit`, change).then(
      json => {
        console.log(json);
      },
      error => {
        console.log(error);
        forceUpdate();
      }
    );
  }

  const handleDelete = id => {
    deleteReq(`/image/${id}/delete`).then(
      json => {
        console.log(json);
        images.delete(id);
        forceUpdate();
      },
      error => {
        console.log(error);
        forceUpdate();
      }
    );
  }

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
      <ImageGallery
        metadata={images}
        onEdit={handleEdit}
        onDelete={handleDelete}
      />
    </React.Fragment >
  )
}

export default Home;
