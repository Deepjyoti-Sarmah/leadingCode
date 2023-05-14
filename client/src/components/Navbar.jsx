import React from 'react'

const Navbar = () => {
    return (
        <div className='bg-black sticky top-0'>
            <header className='relative bg-black'>
                <nav aria-label='Top' className='mx-auto px-4 sm:px-6 lg:px-8'>
                    <div className='flex justify-between h-16 items-center text-white'>
                        {/* logo */}
                        <div className='ml-4 flex lg:ml-0 '>
                            <a href="/">
                                <span >LeetCode</span>
                                {/* <img src="" alt="" /> */}
                            </a>
                        </div>

                        {/* Home button */}
                        {/* <div className='mx-2 lg:ml-8 lg:block lg: self-stretch'>
                            <div className='flex h-full space-x-8'>
                                <a href="/" className='flex items-center text-sm font-medium text-gray-100 hover:text-gray-200 '>
                                    Home
                                </a>
                            </div>
                        </div> */}

                        {/* Problems */}
                        <div className='mx-2 lg:ml-8 lg:block lg: self-stretch'>
                            <div className='flex h-full space-x-8'>
                                <a href="/problemset/all/" className='flex items-center text-sm font-medium text-gray-300 hover:text-gray-100 '>
                                    Problems
                                </a>
                            </div>
                        </div>


                        {/* <div className='ml-auto flex items-center'>
                            <div className='lg:flex lg:flex-1 lg:items-center lg:justify-end lg:space-x-6'>
                                <a href="/login" className='text-sm font-medium text-gray-100 hover:text-gray-200'>Sign in</a>
                            </div>
                        </div> */}
                        <div className="relative ml-4 flex items-center space-x-4 ">
                            <div className="flex items-center">
                                <a
                                    className="text-label-2 text-gray-300 hover:text-gray-100 hidden lg:flex"
                                    href="/signup">
                                    Register
                                </a>
                                <span
                                    className="mx-3 hidden lg:inline-block text-label-3 dark:text-dark-label-3">
                                    or
                                </span>
                                <a
                                    className="text-label-2 text-gray-300 hover:text-gray-100 "
                                    id="navbar_sign_in_button"
                                    href="/login">
                                    Sign in
                                </a>
                            </div>
                        </div>

                    </div>

                </nav>
            </header>
        </div>
    )
}

export default Navbar;