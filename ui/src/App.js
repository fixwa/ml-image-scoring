import './App.css';
import React, {useState} from "react";
import axios from 'axios';

function App() {

    const [file, setFile] = useState('')
    const [serverResponse, setServerResponse] = useState('')

    function handleChange(event) {
        setFile(event.target.files[0])
    }

    function handleSubmit(event) {
        event.preventDefault()
        const url = 'http://localhost:9999/upload';
        const formData = new FormData();
        formData.append('file', file);
        formData.append('fileName', file.name);
        const config = {
            withCredentials: false,
            headers: {
                'content-type': 'multipart/form-data',
            },
        };
        axios.post(url, formData, config).then((response) => {
            console.log(response.data);
            setServerResponse(response.data)
        });

    }

    return (
        <div className="App">
            <form onSubmit={handleSubmit}>
                <h1>Upload a file to analyse...</h1>
                <input type="file" onChange={handleChange}/>
                <button type="submit">Upload</button>
            </form>

            <h3>{serverResponse}</h3>
        </div>
    );
}

export default App;
