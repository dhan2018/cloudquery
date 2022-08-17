import { INTEGRATION_WIDTH } from "./constants";

const GCPLogo = ({ width = INTEGRATION_WIDTH, className }: { width?: number, className?: string }) => (
    <svg
        viewBox="0 0 100 100"
        xmlns="http://www.w3.org/2000/svg"
        xmlSpace="preserve"
        style={{
            fillRule: "evenodd",
            clipRule: "evenodd",
            strokeLinejoin: "round",
            strokeMiterlimit: 2,
        }}
        className={className || "dark:text-white text-gray-900"}
        width={width}
    >
        <path
            d="m66.33 28.882-7.663 7.663a13.734 13.734 0 0 1 5.023 9.413v2.514c.043 0 .083-.014.127-.014 3.87 0 7.02 3.15 7.02 7.02 0 3.87-3.15 7.022-7.02 7.022-.044 0-.084-.012-.127-.013v.013H50.003v10.833H63.69v-.006c.043 0 .083.006.127.006 9.845 0 17.853-8.01 17.853-17.853 0-6.208-3.19-11.682-8.012-14.882a24.588 24.588 0 0 0-7.328-11.716ZM33.333 61.878l-7.85 7.85c2.989 2.252 6.689 3.605 10.704 3.605h13.816V62.5H36.187a6.92 6.92 0 0 1-2.854-.622Zm16.67-39.251c-11.271 0-20.795 7.625-23.685 17.989-4.806 3.201-7.982 8.667-7.982 14.867 0 5.819 2.819 10.984 7.147 14.245l7.85-7.85c-2.446-1.096-4.163-3.546-4.163-6.395 0-3.866 3.15-7.016 7.017-7.016 2.838 0 5.296 1.705 6.401 4.156l7.862-7.861A17.867 17.867 0 0 0 39.802 38a13.72 13.72 0 0 1 18.865-1.455l7.663-7.663c-4.347-3.877-10.058-6.255-16.327-6.255ZM39.802 38a17.887 17.887 0 0 0-3.617-.367 17.746 17.746 0 0 0-9.867 2.983 24.52 24.52 0 0 0-.898 6.594c0 .388.04.767.058 1.15h10.834c-.032-.38-.059-.762-.059-1.15 0-3.539 1.344-6.77 3.549-9.21Z"
            style={{
                fillRule: "nonzero",
            }}
            transform="translate(-28.595 -27.439) scale(1.57195)"
            fill="currentColor"
        />
        <path
            d="M73.988 40.598a24.588 24.588 0 0 0-7.328-11.716l-7.663 7.663a13.734 13.734 0 0 1 5.023 9.413v2.514c.043 0 .083-.014.127-.014 3.87 0 7.02 3.15 7.02 7.02 0 3.87-3.15 7.022-7.02 7.022-.044 0-.084-.012-.127-.013v.013H50.333v10.833H64.02v-.006c.043 0 .083.006.127.006C73.992 73.333 82 65.323 82 55.48c0-6.208-3.19-11.682-8.012-14.882Z"
            style={{
                fill: "#8a8a8a",
                fillOpacity: 0.49,
                fillRule: "nonzero",
            }}
            transform="translate(-29.114 -27.439) scale(1.57195)"
        />
        <path
            d="M50.333 62.5H36.517a6.92 6.92 0 0 1-2.854-.622l-7.85 7.85c2.989 2.252 6.689 3.605 10.704 3.605h13.816V62.5Z"
            style={{
                fill: "#797979",
                fillOpacity: 0.76,
                fillRule: "nonzero",
            }}
            transform="translate(-29.114 -27.439) scale(1.57195)"
        />
        <path
            d="M50.003 22.627c-11.271 0-20.795 7.625-23.685 17.989a17.746 17.746 0 0 1 9.867-2.983c1.239 0 2.449.126 3.617.367a13.72 13.72 0 0 1 18.865-1.455l7.663-7.663c-4.347-3.877-10.058-6.255-16.327-6.255Z"
            style={{
                fill: "#7f7f7f",
                fillOpacity: 0.44,
                fillRule: "nonzero",
            }}
            transform="translate(-28.595 -27.439) scale(1.57195)"
        />
    </svg>
)

export default GCPLogo