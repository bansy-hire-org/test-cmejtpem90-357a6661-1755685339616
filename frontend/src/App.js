import React, { useState, useEffect } from 'react';
import './App.css';

function App() {
  const [licenses, setLicenses] = useState([]);
  const [searchTerm, setSearchTerm] = useState('');

  useEffect(() => {
    const fetchData = async () => {
      const response = await fetch('/api/licenses');
      const data = await response.json();
      setLicenses(data);
    };

    fetchData();
  }, []);

  const handleSearch = (event) => {
    setSearchTerm(event.target.value);
  };

  const filteredLicenses = licenses.filter((license) =>
    license.name.includes(searchTerm)
  );

  return (
    <div className="App">
      <h1>IP License Management</h1>
      <input
        type="text"
        placeholder="Search licenses..."
        value={searchTerm}
        onChange={handleSearch}
      />
      <ul>
        {filteredLicenses.map((license) => (
          <li key={license.id}>{license.name}</li>
        ))}
      </ul>
    </div>
  );
}

export default App;
