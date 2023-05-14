import React from 'react'
import { useParams } from 'react-router-dom';
import problems from '../../data/problemList';
import problemStatements from '../../data/problemStatements';
import ProblemDescription from './ProblemDescription';

const SingleProblem = () => {
    // Get _id from url path
    const id = Number(useParams()._id);

    // Get title, diff from database
    const problem = problems.filter((item) => item._id === id)[0];

    // Get Statement, Examples from database
    const problemDescription = problemStatements.filter((item) => item._id === id)[0];

    return (
        <ProblemDescription
            id={id}
            title={problem.title}
            examples={problemDescription.examples}
            statement={problemDescription.statement}
        />
    );
};

export default SingleProblem;
