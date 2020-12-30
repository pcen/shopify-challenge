import '../styles/images.css';

// ImagesPreview component takes a Map images and returns a grid with a
// preview of each image.
const ImagesPreview = props => {
  const { previewImages } = props;

  if (previewImages.length === 0) {
    return null;
  }

  let previews = [];
  for (let image of previewImages.values()) {
    previews.push(<img className='image-preview' src={image.source} alt={image.name} key={image.name} />);
  }

  return (
    <div className='image-preview-container'>
      {previews}
    </div>
  )
}

export default ImagesPreview;
