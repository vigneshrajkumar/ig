import React from 'react'

import {Link} from 'react-router-dom'

class Nav extends React.Component{
    render(){
        return(
            <div className="nav">
                <div> <Link path="/" >Home</Link>  </div>
                <div> <input name="search" placeholder="Search.." /></div>
                <div> <Link path="/post" >Post</Link>  </div>
                <div> <Link path="/profile" >Profile</Link>  </div>
                <div> <Link path="/logout" >Logout</Link>  </div>
            </div>
        )
    }
}

export default Nav