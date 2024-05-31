//Login.js
import React, {useState} from 'react';

const Login = ({auth, handleLogin, handleLogout, userInfo}) => {
    const [quote, setQuote] = useState('');
    const [error, setError] = useState('');
    const [loading, setLoading] = useState(false);

    const randomString = () => {
        setLoading(true);
        fetch('/random/string', { // or use 'http://localhost:5000/random/string'
            headers: {
                'Authorization': `Bearer ${userInfo.access_token}`
            }
        })
            .then(response => {
                if (response.ok) {
                    return response.json();
                } else if (response.status === 401) {
                    throw new Error('Authentication required, please log in.');
                } else if (response.status === 403) {
                    throw new Error('Access forbidden, you do not have necessary permissions.');
                } else {
                    throw new Error('Error while fetching data');
                }
            })
            .then(data => {
                setQuote(data.quote);
                setError('');
                setLoading(false);
            })
            .catch((error) => {
                setError(error.message);
                setLoading(false);
            });
    };

    const randomInteger = () => {
        setLoading(true);
        fetch('/random/integer', {
            headers: {
                'Authorization': `Bearer ${userInfo.access_token}`
            }
        })
            .then(response => {
                if (response.ok) {
                    return response.json();
                } else if (response.status === 401) {
                    throw new Error('Authentication required, please log in.');
                } else if (response.status === 403) {
                    throw new Error('Access forbidden, you do not have necessary permissions.');
                } else {
                    throw new Error('Error while fetching data');
                }
            })
            .then(data => {
                setQuote(data.quote);
                setError('');
                setLoading(false);
            })
            .catch((error) => {
                setError(error.message);
                setLoading(false);
            });
    };

    if (auth === null) {
        return <div>Loading...</div>;
    }

    if (auth === false) {
        return (
            <div>
                <h1>Welcome!</h1>
                <button onClick={handleLogin}>Please log in.</button>
            </div>
        );
    }

    if (auth === true && userInfo) {
        return (
            <div>
                <h1>Welcome, {userInfo.profile.name}!</h1>
                {/* <h2>Your access token: {userInfo.access_token}</h2> */}
                <button onClick={handleLogout}>Log out</button>
                <button onClick={randomString}>Generate Random String</button>
                <button onClick={randomInteger}>Generate Random Integer</button>
                {loading && <p>Loading ...</p>}
                {error && <p>{error}</p>}
                {quote && <p>Random Result: <b>{quote}</b>.</p>}
                {quote && <p>Name <b>{userInfo.profile.name}</b></p>}
                {quote && <p>Email <b>{userInfo.profile.email}</b>.</p>}
            </div>
        );
    }

    return <div>Loading...</div>;
};

export default Login;
