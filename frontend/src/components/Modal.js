import Popup from 'reactjs-popup';

import '../styles/modal.css';

const Modal = props => {
  const { trigger, onClose, content } = props;

  return (
    <Popup
      trigger={trigger}
      onClose={onClose}
      modal
      position='top center'
      closeOnDocumentClick={false}
      closeOnEscape={true}
    >
      {close => (
        <div className='modal-body'>
          <button className='close' onClick={close}>&times;</button>
          {content}
        </div>
      )}
    </Popup>
  )
}

export default Modal;
