import '../styles/gallery.css';

const ImageGallery = props => {
  const { content } = props;

  return (
    <div className='image-gallery'>
      {content}
    </div>
  )
}

export default ImageGallery;
