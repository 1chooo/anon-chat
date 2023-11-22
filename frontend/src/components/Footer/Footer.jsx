import React from 'react';
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
        <footer className="Footer">
            <p> © 2023 Copyright: Hugo ChunHo Lin</p>
        </footer>
    );
}
