import React, { useEffect, useState } from 'react'
import { Link } from "react-router-dom";

export const ProblemSet = () => {

    const [problems, setProblems] = useState([]);

    const init = async () => {
        const response = await fetch("http://localhost:3000/problems", {
            method: "GET",
        });

        const json = await response.json();
        setProblems(json.problems)
    }

    useEffect(() => {
        init()
    }, []);

    return (
        <div className="overflow-x-auto max-w-7xl mx-auto my-16">
            <table className="min-w-full divide-y-2 divide-gray-200 text-sm">
                <thead className="ltr:text-left rtl:text-right">
                    <tr>
                        <th className="whitespace-nowrap px-4 py-2 font-medium text-gray-900">
                            Title
                        </th>
                        <th className="whitespace-nowrap px-4 py-2 font-medium text-gray-900">
                            Difficulty
                        </th>
                        <th className="whitespace-nowrap px-4 py-2 font-medium text-gray-900">
                            Acceptance
                        </th>
                        <th className="px-4 py-2"></th>
                    </tr>
                </thead>

                <tbody className="divide-y divide-gray-200">
                    {problems.map((item, id) => <tr key={id}>
                        <td className="whitespace-nowrap px-4 py-2 font-medium text-gray-900">
                            <Link to={`/problems/:${item.problemId}`}>{item.problemId}. {item.title}</Link>
                        </td>
                        <td className="whitespace-nowrap px-4 py-2 text-gray-700 text-center">
                            {item.difficulty}
                        </td>
                        <td className="whitespace-nowrap px-4 py-2 text-gray-700 text-center">
                            {item.acceptance}
                        </td>
                        <td className="whitespace-nowrap px-4 py-2">
                            <Link
                                to={`/problems/:${item.problemId}`}
                                className="inline-block rounded bg-indigo-600 px-4 py-2 text-xs font-medium text-white hover:bg-indigo-700"
                            >
                                View
                            </Link>
                        </td>
                    </tr>
                    )}

                </tbody>
            </table>
        </div>
    )
}

// export default ProblemSet;