import React, { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom';
import ProblemDescription from './ProblemDescription';
import {backendURL} from "../constant"


const SingleProblem = () => {

    const [CodeSeg, setCodeSeg] = useState("");
    const { pid } = useParams() ;
    const cleanId = pid.substring(1) ;
    const [problem, setProblem] = useState(null);
    // const [submission, setSubmission] = useState("");
    // console.log(cleanId);
    // const [allsubmission, setAllsubmission] = useState(null);


    const init = async() => {
        const response = await fetch(`${backendURL}/problems/`+ cleanId, {
            method: "GET",
        });
    

        const json = await response.json();
        setProblem(json.problem)

        // getAllSubmission();
    }

    // const getAllSubmission = async () => {
    //     const allsubmissionResponse = await fetch(`${backendURL}/submissions/`+cleanId, {
    //         method: "GET",
    //         headers: {
    //             "authorization": localStorage.getItem("token")
    //         },
    //     });
    //     const allsubmissionJson = await allsubmissionResponse.json();
    //     setAllsubmission(allsubmissionJson.submissions);
    // }

    useEffect(() => {
        init();
    },[])

    

    return (
        problem? (
            <ProblemDescription
                id={problem.problemId}
                title={problem.title}
                statement={problem.description}
                exampleIn={problem.exampleIn}
                exampleOut={problem.exampleOut}
                cleanId = {cleanId}
                // allSubmission = {allsubmission}
            />
        ):(<div>The searched Question Doesn't exist</div>)
    );
};

export default SingleProblem;
