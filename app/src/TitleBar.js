import React from 'react';
import dateformat from 'dateformat';
import './TitleBar.scss';

class TitleBar extends React.Component {

    constructor(props) {
        super(props);
    }

    style() { return { backgroundColor: '#762124' } }

    time() { return dateformat(new Date(), 'HH:MM:ss') }

    render() {
        return (
            <nav className="navbar navbar-expand navbar-dark" style={this.style()}>
                <a className='navbar-brand mr-0'>
                    <img src={'https://www.aueb.gr/press/logos/1_AUEB-pantone-HR.jpg'}
                         height="44" />
                </a>
                <nav className="nav navbar mr-auto">
                    <a className="navbar-brand">
                        AUEB CSLabs
                    </a>
                </nav>
                <a className="navbar-brand">
                    {this.time()}
                </a>
            </nav>
        );
    }
}

export default TitleBar;
