import React, { useState } from 'react'

const CooForm = () => {

    const [msg, setMsg] = useState("")

  return (<>
    <table className='student-form' style={{marginTop :"10px"}}>
            <tr>
                <td>Departments</td>
                <td>
                    <input type='text' />
                </td>
            </tr>
            <tr>
                <td>Rooms</td>
                <td>
                    <input type='text' />
                </td>
            </tr>
            <tr>
                <td>Date</td>
                <td><input type="date" name="" id="date" /></td>
            </tr>
            <tr>
                <td></td>
                <td><div className='stu-form-button' onClick={() => {setMsg("Under Development")}}>Submit</div></td>
            </tr>
        </table>

        {msg !== "" && <h3 style={{marginTop:"10px", color:"green"}}>{msg}</h3>}

</>
  )
}

export default CooForm