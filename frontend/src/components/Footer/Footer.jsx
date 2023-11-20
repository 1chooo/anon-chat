import React from 'react';
import {
    MDBFooter,
    MDBContainer,
    MDBIcon,
    MDBBtn
} from 'mdb-react-ui-kit';
import './Footer.scss';

export default function Footer() {
    const socialMediaLinks = [
        { icon: ['fab', 'facebook-f'], link: 'https://www.facebook.com/' },
        { icon: ['fab', 'twitter'], link: 'https://www.twitter.com/' },
        { icon: ['fab', 'google'], link: 'https://www.google.com/' },
        { icon: ['fab', 'instagram'], link: 'https://www.instagram.com/' },
        { icon: ['fab', 'linkedin'], link: 'https://www.linkedin.com/' },
        { icon: ['fab', 'github'], link: 'https://www.github.com/' }
    ];

    return (
        <MDBFooter className='custom-footer'>
            <MDBContainer className='pt-4'>
                <section className='mb-4'>
                    {socialMediaLinks.map((item, index) => (
                        <MDBBtn
                            key={index}
                            rippleColor='dark'
                            color='link'
                            floating
                            size='lg'
                            className='text-dark m-1 social-media-btn'
                            href={item.link}
                            role='button'
                            target='_blank'
                            rel='noopener noreferrer'
                        >
                            <MDBIcon fab icon={item.icon} />
                        </MDBBtn>
                    ))}
                </section>
            </MDBContainer>

            <div className='custom-footer-text'>
                © 2023 Copyright:
                <a className='text-dark' href='https://mdbootstrap.com/' target='_blank' rel='noopener noreferrer'>
                    MDBootstrap.com
                </a>
            </div>
        </MDBFooter>
    );
}
