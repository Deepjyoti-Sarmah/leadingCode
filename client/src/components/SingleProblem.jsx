import React, { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom';
import problems from '../../data/problemList';
import problemStatements from '../../data/problemStatements';
import ProblemDescription from './ProblemDescription';

const SingleProblem = () => {

    const [CodeSeg, setCodeSeg] = useState("");
    const { pid } = useParams() ;
    const cleanId = pid.substring(1) ;
    const [problem, setProblem] = useState(null);
    // const [submission, setSubmission] = useState("");
    console.log(cleanId);

    const init = async() => {
        const response = await fetch("http://localhost:3000/problems/"+ cleanId, {
            method: "GET",
        });
    

        const json = await response.json();
        setProblem(json.problem)
    }

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
            />
        ):(<div>The searched Question Doesn't exist</div>)
    );
};

export default SingleProblem;
