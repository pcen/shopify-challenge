import { useState } from "react";

const SearchBar = props => {
  const [previous, setPrevious] = useState(null);
  const [query, setQuery] = useState('');

  const { onSubmit } = props;

  const handleChange = value => {
    setQuery(value);
  }

  const handleKeyPress = key => {
    if (key === 'Enter' && query !== previous) {
      onSubmit(query);
      setPrevious(query);
    }
  }

  return (
    <input
      type='text'
      placeholder='Search Images...'
      style={{
        padding: '2px 3px 2px 3px',
        minWidth: '300px',
      }}
      onChange={event => handleChange(event.target.value)}
      onKeyPress={event => handleKeyPress(event.key)}
    >
    </input>
  )
}

export default SearchBar;
