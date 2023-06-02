import React, { useEffect } from 'react'
import { Editor } from "@monaco-editor/react";
import { Listbox, Transition } from "@headlessui/react";
import { Fragment, useState } from 'react';
import { CheckIcon, ChevronUpDownIcon } from "@heroicons/react/24/outline";
import defaultCodeSnippets from "./defaultCodeSnippets"
import { useParams } from 'react-router-dom';
import { backendURL } from '../constant';

const corsOptions = {
    origin: `${backendURL}`, // Replace with your backend domain
    methods: ['GET', 'POST', 'PUT', 'DELETE'],
    allowedHeaders: ['Content-Type', 'Authorization'],
};

const ProblemDescription = (props) => {
    const { id, title, statement, exampleIn, exampleOut , cleanId} = props;
    const [solutionLanguage, setSolutionLanguage] = useState("cpp");

    const [submission, setSubmission] = useState("");
    const [allsubmission, setAllsubmission] = useState(null);

    const [showTable, setShowTable] = useState(false);

    const handleButtonClick = async () => {
        setShowTable(prevState => !prevState);
        
        if (!showTable) {
            // const allsubmissionResponse = await fetch(`${backendURL}/submissions/`+cleanId, {
            //     method: "GET",
            //     headers: {
            //         "authorization": localStorage.getItem("token")
            //     },
            // });
            // const allsubmissionJson = await allsubmissionResponse.json();
            // setAllsubmission(allsubmissionJson.submissions);
            getAllSubmission();
        }
    }

    const getAllSubmission = async () => {
        const allsubmissionResponse = await fetch(`${backendURL}/submissions/`+cleanId, {
            method: "GET",
            headers: {
                "authorization": localStorage.getItem("token")
            },
        });
        const allsubmissionJson = await allsubmissionResponse.json();
        setAllsubmission(allsubmissionJson.submissions);
        // console.log(allsubmissionResponse);
    }

    function handleEditorChange(value, event) {
        setSubmission(value);
    }

    // const handleKey = (event) => {
    //     if (event.key == "Tab"){
    //         event.preventDefault();
    //         const {selectionStart, selectionEnd, value} = event.target ;
    //         const val = value.substring(0, selectionStart)+ "\t" + value.substring(selectionStart);
    //         event.target.value = val;
    //         event.target.selectionStart = event.target.selectionEnd = selectionStart+1;
    //     }
    //     setCodeSeg(event.value);
    // }

    useEffect(() => {
    getAllSubmission();
    },[])

    return (
        <div className="text-gray-100 min-h-screen bg-gray-950">
            <div className="mx-auto px-6 py-4">
                <div className="grid grid-cols-2 gap-6 items-start">
                    <div className="grid grid-cols-1">
                        <h1 className="text-xl font-bold">
                            {id}. {title}
                        </h1>
                        <div className="mt-4 text-base font-normal">{statement}</div>
                        {/* {examples.map((example, idx) => (
                            <div>
                                <div className="px-1 mt-6 mb-1 font-medium">
                                    Example {idx + 1}:
                                </div> */}
                                <div className="border-gray-700 bg-slate-600 rounded-2xl px-3 py-2 mt-6 mb-1">
                                    <div>
                                        <div className="font-bold">Input :</div>
                                        <div className="font-mono">{exampleIn}</div>
                                    </div>
                                    <div>
                                        <div className="font-bold">Output :</div>
                                        <div className="font-mono">{exampleOut}</div>
                                    </div>
                                </div>

                                <div className= "flex justify-start w-full mt-auto max-w-xs px-2 py-2 text-center items-center">
                                    <button
                                        className="flex w-[1/4] justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
                                        onClick={handleButtonClick}
                                    >View Submissions</button>
                                </div>
                                    {
                                        showTable && allsubmission && (
                                            <div className='px-2 py-2 '>
                                                <table className=" divide-y-2 divide-gray-200 text-sm">
                                                    <tbody>
                                                        <tr>
                                                            <th className='whitespace-nowrap px-2 py-2 font-medium text-white-900'>Problem ID</th>
                                                            <th className='whitespace-nowrap px-2 py-2 font-medium text-white-900'>Status</th>
                                                        </tr>

                                                        {allsubmission.map((prob, index) => (
                                                            <tr key= {index}>
                                                                <td className="px-2 py-2 font-medium ">{prob.problemId}</td>
                                                                <td className="px-2 py-2 font-medium ">{prob.status}</td>
                                                            </tr>
                                                        ))}
                                                    </tbody>
                                                </table>
                                            </div>
                                        )
                                    }
                            {/* </div> */}
                        {/* ))} */}
                    </div>
                    
                    <div className="grid grid-cols-1 border-orange-400">
                        <div className="">
                            <Editor
                                height="80vh"
                                language={solutionLanguage}
                                defaultValue={defaultCodeSnippets.javascript}
                                value={defaultCodeSnippets[solutionLanguage]}
                                theme="vs-dark"
                                onChange={handleEditorChange}
                                onKeyDown= {(event) => handleKey(event)}
                            />
                        </div>
                        <div className="flex justify-between text-black mt-auto">
                            <div className='flex justify-start w-full mt-auto max-w-xs text-center items-center'>
                                <button 
                                    type="submit"
                                    className="flex align-middle items-center w-full rounded-md bg-green-600 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-green-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-green-600"
                                    onClick={ async () => {
                                        const response = await fetch(`${backendURL}/submission`, {
                                            method: "POST",
                                            headers: {
                                                "authorization": localStorage.getItem("token")
                                            },
                                            body: JSON.stringify({
                                                probelemId: cleanId,
                                                submission: submission
                                            })
                                        });
                                        const json = await response.json();
                                        console.log(json);
                                    }}
                                    >Submit
                                </button>
                            </div>
                            <div className="flex justify-end text-black w-full mt-auto">
                                <div className="w-full mt-auto max-w-xs">
                                    <Listbox value={solutionLanguage} onChange={setSolutionLanguage}>
                                        <div className="w-full items-right justify-right border-orange-700">
                                            <div className="relative py-1 px-1">
                                                <Listbox.Button className="relative w-full cursor-default rounded-lg bg-gray-500 py-2 pl-3 pr-10 text-left shadow-md focus:outline-none focus-visible:border-indigo-500 focus-visible:ring-2 focus-visible:ring-white focus-visible:ring-opacity-75 focus-visible:ring-offset-2 focus-visible:ring-offset-orange-300 sm:text-sm">
                                                    <span className="block truncate">{solutionLanguage}</span>
                                                    <span className="pointer-events-none absolute inset-y-0 right-0 flex items-center pr-2">
                                                        <ChevronUpDownIcon
                                                            className="h-5 w-5 text-gray-400"
                                                            aria-hidden="true"
                                                        />
                                                    </span>
                                                </Listbox.Button>
                                                <Transition
                                                    as={Fragment}
                                                    leave="transition ease-in duration-100"
                                                    leaveFrom="opacity-100"
                                                    leaveTo="opacity-0"
                                                    className="bottom-full"
                                                >
                                                    <Listbox.Options className="absolute mb-1 w-full max-h-60 overflow-auto rounded-md bg-white py-1 text-base shadow-md ring-1 ring-black ring-opacity-5 focus:outline-none sm:text-sm">
                                                        {Object.keys(defaultCodeSnippets).map(
                                                            (language, idx) => (
                                                                <Listbox.Option
                                                                    key={idx}
                                                                    className={({ active }) =>
                                                                        `relative cursor-default select-none py-2 pl-10 pr-4 ${active
                                                                            ? "bg-amber-100 text-amber-900"
                                                                            : "text-gray-900"
                                                                        }`
                                                                    }
                                                                    value={language}
                                                                >
                                                                    {({ selected }) => (
                                                                        <>
                                                                            <span
                                                                                className={`block truncate ${selected ? "font-medium" : "font-normal"
                                                                                    }`}
                                                                            >
                                                                                {language}
                                                                            </span>
                                                                            {selected ? (
                                                                                <span className="absolute inset-y-0 left-0 flex items-center pl-3 text-amber-600">
                                                                                    <CheckIcon
                                                                                        className="h-5 w-5"
                                                                                        aria-hidden="true"
                                                                                    />
                                                                                </span>
                                                                            ) : null}
                                                                        </>
                                                                    )}
                                                                </Listbox.Option>
                                                            )
                                                        )}
                                                    </Listbox.Options>
                                                </Transition>
                                            </div>
                                        </div>
                                    </Listbox>
                                </div>
                            </div>
                        </div>

                    </div>
                </div>
            </div>
        </div>
    )
}

export default ProblemDescription;