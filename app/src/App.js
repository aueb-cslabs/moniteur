import React from 'react';
import './App.scss';
import TitleBar from './TitleBar'
import Schedule from './widgets/Schedule'

function App() {
  return (
      <div>
        <TitleBar/>
        <div className="WidgetContainer p-4">
          <Schedule height={2} width={2}/>
        </div>
      </div>
  );
}

export default App;
