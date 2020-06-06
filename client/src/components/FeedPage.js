import React from 'react'

import Post from "./Post"
import Nav from "./Nav"

class FeedPage extends React.Component {
    render() {
        return (
            <div className="app">
                <Nav />
                <Post />
                <Post />
            </div>
        )
    }
}

export default FeedPage