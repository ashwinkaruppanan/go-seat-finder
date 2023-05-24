import React, { useState } from 'react'
import StudentForm from './StudentForm'
import CooForm from './CooForm'

const App = () => {

  const [menu, setMenu] = useState(true)

  return (<><center>
    <h2>Exam Seat Finder</h2>
    <div className='tab'>
      <div  className={menu ? "tab-c tab-t" : "tab-c tab-f"} onClick={() => setMenu(true)}>Student</div>
      <div className={!menu ? "tab-c tab-t" : "tab-c tab-f"} onClick={() => setMenu(false)}>Exam Coordinator</div>
    </div>
    {menu ? <StudentForm /> : <CooForm />}
    </center>
    </>
  )
}

export default App