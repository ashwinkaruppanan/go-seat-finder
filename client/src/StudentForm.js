import React, { useState } from 'react'
import axios from "axios"


const StudentForm = () => {

    const [res , setRes] = useState({
        status : 20,
        room_number : 0,
        seat_number : 0,
        date : ""
    })

    async function handleClick() {
        let rollNumber = document.getElementById("rollNumber").value
        let data = {
            "roll_number" : rollNumber
        }        
        try{
            const response = await axios.post("http://localhost:8080/findseat", JSON.stringify(data));
            const date = new Date(response.data.date * 1000);
            const dateStringOnly = date.toLocaleDateString();
            setRes(pres => ({
                ...pres, 
                status : 200,
                room_number : response.data.room_number,
                seat_number : response.data.seat_number,
                date : dateStringOnly
            
            }))
            
        } catch (error) {
            if (error.response.status === 500) {
                setRes(pres => ({
                    ...pres, 
                    status : 500
                })) 
            }
        }
    }

  return (<>
        <p>sample rollnumber - 181CA001, 181CS001, 181IT001, 181DS001</p>
        <table className='student-form' style={{marginTop :"10px"}}>
            <tr>
                <td>Roll Number</td>
                <td><input type="text" name="" id="rollNumber" /></td>
            </tr>
            {/* <tr>
                <td>Date</td>
                <td><input type="date" name="" id="date" /></td>
            </tr> */}
            <tr>
                <td></td>
                <td><div className='stu-form-button' onClick={() => {handleClick()}}>Submit</div></td>
            </tr>
        </table>
    

    {res.status === 500 && <h2 style={{color : "red", marginTop: "20px"}}>No records found</h2>}

    {res.status === 200 && <><table style={{marginTop:"20px"}} className='td-border'>
        <tr>
            <td className='td-padd'>
                Room Number
            </td>
            <td className='td-padd'>
                {res.room_number}
            </td>
        </tr>
        <tr>
            <td className='td-padd'>
                Seat Number
            </td>
            <td className='td-padd'>
                {res.seat_number}
            </td>
        </tr>
        </table>
        
        <h3 style={{marginTop :"5px"}}>{res.date}</h3>
        </>
        }
    
    </>
  )
}

export default StudentForm