import React, { useState } from 'react';
import Home from './pages/Home/Home';

import './styles/styles.css';

const App = () => {
    const [data, setData] = useState({
        rocketMass: '',
        fuelMass: '',
        uFunction: '',
        fFunction: ''
    });
    const handleDataChange = (data) => {
        setData(data);
        console.log(data);
    }
    return <Home data={data} handleChangeData={handleDataChange} />
}

export default App;