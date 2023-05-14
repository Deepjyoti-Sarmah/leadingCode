import React from 'react'

const Home = () => {
    return (
        <div className='flex flex-col justify-center align-middle items-center min-h-screen bg-gray-950 text-white'>
            <h1 className='font-extrabold text-3xl mx-3 my-2 p-3'>
                <span>
                    Hello Welcome
                </span>
                <span className="text-xl px-2 py-2 font-extralight">
                    to
                </span>
                <span className='text-4xl px-1 '>
                    LeetCode
                </span>
                <span className=' flex flex-col text-2xl font-light p-3 mb-2'>
                    a platform to solve coding
                    <span className='flex flex-col font-extrabold text-3xl text-blue-500'>
                        <a href="/problems/all">{" "}Problems</a>
                    </span>
                    {/* </h2> */}
                </span>
            </h1>
            {/* <h2 className='font-bold'>to</h2>
            <h1 className='text-3xl font-extrabold mx-3 my-2 p-3'>Leetcode</h1> */}
            {/* <h2 className='font-bold'>a platform to solve coding
                <span className='font-extrabold text-2xl text-blue-500'>
                    <a href="/problemset/all">{" "}Problems</a>
                </span>
            </h2> */}
        </div>
    )
}

export default Home