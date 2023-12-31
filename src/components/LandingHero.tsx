import gsap from "gsap";
import { useLayoutEffect, useRef } from "react";

function LandingHero() {
  const root = useRef<SVGSVGElement>(null);
  const tl = useRef<GSAPTimeline>();

  useLayoutEffect(() => {
    gsap.context(() => {
      tl.current = gsap
        .timeline({ repeat: 0 })
        .set("#svgCalendar", { y: 0 })
        .to("#svgMouse", { x: -350, y: -250, duration: 1 })
        .to("#svgMouse", { opacity: 0, duration: 0.35, delay: 0.2 })
        .to("#svgButtonText", { opacity: 0, duration: 0.35 })
        .to("#svgButtonBackground", { attr: { fill: "#19141a" } }, "<")
        .to(
          "#svgButton",
          {
            scale: 0,
            transformOrigin: "50% 50%",
            duration: 1,
          },
          "<"
        )
        .to("#svgButton", { opacity: 0, duration: 0.5 }, "<")
        .to("#svgCalendar", { scale: 1, transformOrigin: "50% 50%" }, "<")
        .to("#svgCalendar line", { opacity: 1, duration: 0.5 })
        .to(
          "#svgCalendar g",
          {
            opacity: 1,
            duration: 0.5,
            stagger: { grid: [2, 10], from: "start", axis: "x", amount: 1.5 },
          },
          "<"
        );
    }, root);
  }, []);

  return (
    <svg
      width="850"
      height="612"
      viewBox="0 0 850 612"
      style={{
        width: "100%",
        height: "100%",
      }}
      className="max-md:max-h-[40vh]"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
      ref={root}
    >
      <g clipPath="url(#clip0_41_728)" id="svgCalendar" scale={0}>
        <rect
          width="848"
          height="608"
          rx="16"
          className="fill-[var(--mantine-color-gray-4)] dark:fill-[var(--mantine-color-dark-8)]"
        />
        <line
          x1="175.3"
          y1="-49.9986"
          x2="173.3"
          y2="658.001"
          stroke="black"
          opacity={0}
        />
        <line
          x1="342.1"
          y1="-49.9986"
          x2="340.1"
          y2="658.001"
          stroke="black"
          opacity={0}
        />
        <line
          x1="508.9"
          y1="-49.9986"
          x2="506.9"
          y2="658.001"
          stroke="black"
          opacity={0}
        />
        <line
          x1="675.7"
          y1="-49.9986"
          x2="673.7"
          y2="658.001"
          stroke="black"
          opacity={0}
        />
        <g opacity={0}>
          <rect
            x="18"
            y="241"
            width="145"
            height="168"
            rx="8"
            className="fill-[var(--mantine-primary-color-filled)]"
          />
          <path
            d="M33 258H140"
            className="stroke-[var(--mantine-color-myColor-4)] dark:stroke-[var(--mantine-color-myColor-6)]"
            strokeWidth="8"
            strokeLinecap="round"
          />
          <path
            d="M33 278H97"
            className="stroke-[var(--mantine-color-myColor-4)] dark:stroke-[var(--mantine-color-myColor-6)]"
            strokeWidth="8"
            strokeLinecap="round"
          />
          <path
            d="M33 298H111"
            className="stroke-[var(--mantine-color-myColor-4)] dark:stroke-[var(--mantine-color-myColor-6)]"
            strokeWidth="8"
            strokeLinecap="round"
          />
        </g>

        <g opacity={0}>
          <rect
            x="185"
            y="110"
            width="145"
            height="168"
            rx="8"
            className="fill-[var(--mantine-primary-color-filled)]"
          />
          <path
            d="M200 127H307"
            className="stroke-[var(--mantine-color-myColor-4)] dark:stroke-[var(--mantine-color-myColor-6)]"
            strokeWidth="8"
            strokeLinecap="round"
          />
          <path
            d="M200 147H264"
            className="stroke-[var(--mantine-color-myColor-4)] dark:stroke-[var(--mantine-color-myColor-6)]"
            strokeWidth="8"
            strokeLinecap="round"
          />
          <path
            d="M200 167H278"
            className="stroke-[var(--mantine-color-myColor-4)] dark:stroke-[var(--mantine-color-myColor-6)]"
            strokeWidth="8"
            strokeLinecap="round"
          />
        </g>

        <g opacity={0}>
          <rect
            x="185"
            y="341"
            width="145"
            height="109"
            rx="8"
            className="fill-[var(--mantine-primary-color-filled)]"
          />
          <path
            d="M200 358H307"
            className="stroke-[var(--mantine-color-myColor-4)] dark:stroke-[var(--mantine-color-myColor-6)]"
            strokeWidth="8"
            strokeLinecap="round"
          />
          <path
            d="M200 378H264"
            className="stroke-[var(--mantine-color-myColor-4)] dark:stroke-[var(--mantine-color-myColor-6)]"
            strokeWidth="8"
            strokeLinecap="round"
          />
          <path
            d="M200 398H278"
            className="stroke-[var(--mantine-color-myColor-4)] dark:stroke-[var(--mantine-color-myColor-6)]"
            strokeWidth="8"
            strokeLinecap="round"
          />
        </g>

        <g opacity={0}>
          <rect
            x="351"
            y="419"
            width="145"
            height="76"
            rx="8"
            className="fill-[var(--mantine-primary-color-filled)]"
          />
          <path
            d="M366 436H473"
            className="stroke-[var(--mantine-color-myColor-4)] dark:stroke-[var(--mantine-color-myColor-6)]"
            strokeWidth="8"
            strokeLinecap="round"
          />
          <path
            d="M366 456H430"
            className="stroke-[var(--mantine-color-myColor-4)] dark:stroke-[var(--mantine-color-myColor-6)]"
            strokeWidth="8"
            strokeLinecap="round"
          />
          <path
            d="M366 476H444"
            className="stroke-[var(--mantine-color-myColor-4)] dark:stroke-[var(--mantine-color-myColor-6)]"
            strokeWidth="8"
            strokeLinecap="round"
          />
        </g>

        <g opacity={0}>
          <rect
            x="18"
            y="422"
            width="145"
            height="76"
            rx="8"
            className="fill-[var(--mantine-primary-color-filled)]"
          />
          <path
            d="M33 439H140"
            className="stroke-[var(--mantine-color-myColor-4)] dark:stroke-[var(--mantine-color-myColor-6)]"
            strokeWidth="8"
            strokeLinecap="round"
          />
          <path
            d="M33 459H97"
            className="stroke-[var(--mantine-color-myColor-4)] dark:stroke-[var(--mantine-color-myColor-6)]"
            strokeWidth="8"
            strokeLinecap="round"
          />
          <path
            d="M33 479H111"
            className="stroke-[var(--mantine-color-myColor-4)] dark:stroke-[var(--mantine-color-myColor-6)]"
            strokeWidth="8"
            strokeLinecap="round"
          />
        </g>

        <g opacity={0}>
          <rect
            x="351"
            y="194"
            width="145"
            height="168"
            rx="8"
            className="fill-[var(--mantine-primary-color-filled)]"
          />
          <path
            d="M366 211H473"
            className="stroke-[var(--mantine-color-myColor-4)] dark:stroke-[var(--mantine-color-myColor-6)]"
            strokeWidth="8"
            strokeLinecap="round"
          />
          <path
            d="M366 231H430"
            className="stroke-[var(--mantine-color-myColor-4)] dark:stroke-[var(--mantine-color-myColor-6)]"
            strokeWidth="8"
            strokeLinecap="round"
          />
          <path
            d="M366 251H444"
            className="stroke-[var(--mantine-color-myColor-4)] dark:stroke-[var(--mantine-color-myColor-6)]"
            strokeWidth="8"
            strokeLinecap="round"
          />
        </g>

        <g opacity={0}>
          <rect
            x="518"
            y="207"
            width="145"
            height="110"
            rx="8"
            className="fill-[var(--mantine-primary-color-filled)]"
          />
          <path
            d="M533 224H640"
            className="stroke-[var(--mantine-color-myColor-4)] dark:stroke-[var(--mantine-color-myColor-6)]"
            strokeWidth="8"
            strokeLinecap="round"
          />
          <path
            d="M533 244H597"
            className="stroke-[var(--mantine-color-myColor-4)] dark:stroke-[var(--mantine-color-myColor-6)]"
            strokeWidth="8"
            strokeLinecap="round"
          />
          <path
            d="M533 264H611"
            className="stroke-[var(--mantine-color-myColor-4)] dark:stroke-[var(--mantine-color-myColor-6)]"
            strokeWidth="8"
            strokeLinecap="round"
          />
        </g>

        <g opacity={0}>
          <rect
            x="518"
            y="110"
            width="145"
            height="84"
            rx="8"
            className="fill-[var(--mantine-primary-color-filled)]"
          />
          <path
            d="M533 127H640"
            className="stroke-[var(--mantine-color-myColor-4)] dark:stroke-[var(--mantine-color-myColor-6)]"
            strokeWidth="8"
            strokeLinecap="round"
          />
          <path
            d="M533 147H597"
            className="stroke-[var(--mantine-color-myColor-4)] dark:stroke-[var(--mantine-color-myColor-6)]"
            strokeWidth="8"
            strokeLinecap="round"
          />
          <path
            d="M533 167H611"
            className="stroke-[var(--mantine-color-myColor-4)] dark:stroke-[var(--mantine-color-myColor-6)]"
            strokeWidth="8"
            strokeLinecap="round"
          />
        </g>

        <g opacity={0}>
          <rect
            x="685"
            y="192"
            width="145"
            height="106"
            rx="8"
            className="fill-[var(--mantine-primary-color-filled)]"
          />
          <path
            d="M700 209H807"
            className="stroke-[var(--mantine-color-myColor-4)] dark:stroke-[var(--mantine-color-myColor-6)]"
            strokeWidth="8"
            strokeLinecap="round"
          />
          <path
            d="M700 229H764"
            className="stroke-[var(--mantine-color-myColor-4)] dark:stroke-[var(--mantine-color-myColor-6)]"
            strokeWidth="8"
            strokeLinecap="round"
          />
          <path
            d="M700 249H778"
            className="stroke-[var(--mantine-color-myColor-4)] dark:stroke-[var(--mantine-color-myColor-6)]"
            strokeWidth="8"
            strokeLinecap="round"
          />
        </g>

        <g opacity={0}>
          <rect
            x="685"
            y="311"
            width="145"
            height="106"
            rx="8"
            className="fill-[var(--mantine-primary-color-filled)]"
          />
          <path
            d="M700 328H807"
            className="stroke-[var(--mantine-color-myColor-4)] dark:stroke-[var(--mantine-color-myColor-6)]"
            strokeWidth="8"
            strokeLinecap="round"
          />
          <path
            d="M700 348H764"
            className="stroke-[var(--mantine-color-myColor-4)] dark:stroke-[var(--mantine-color-myColor-6)]"
            strokeWidth="8"
            strokeLinecap="round"
          />
          <path
            d="M700 368H778"
            className="stroke-[var(--mantine-color-myColor-4)] dark:stroke-[var(--mantine-color-myColor-6)]"
            strokeWidth="8"
            strokeLinecap="round"
          />
        </g>
      </g>
      <g filter="url(#filter0_d_41_728)" id="svgButton">
        <rect
          x="303"
          y="271"
          width="242"
          height="65"
          rx="16"
          fill="#1E1C21"
          className="fill-[var(--mantine-primary-color-filled)]"
          shapeRendering="crispEdges"
          id="svgButtonBackground"
        />
        <path
          d="M388.091 291.727H390.909V307.136C390.909 308.727 390.534 310.148 389.784 311.398C389.042 312.64 387.992 313.621 386.636 314.341C385.28 315.053 383.689 315.409 381.864 315.409C380.038 315.409 378.447 315.053 377.091 314.341C375.735 313.621 374.682 312.64 373.932 311.398C373.189 310.148 372.818 308.727 372.818 307.136V291.727H375.636V306.909C375.636 308.045 375.886 309.057 376.386 309.943C376.886 310.822 377.598 311.515 378.523 312.023C379.455 312.523 380.568 312.773 381.864 312.773C383.159 312.773 384.273 312.523 385.205 312.023C386.136 311.515 386.848 310.822 387.341 309.943C387.841 309.057 388.091 308.045 388.091 306.909V291.727ZM396.173 321.545V297.545H398.764V300.318H399.082C399.279 300.015 399.552 299.629 399.901 299.159C400.257 298.682 400.764 298.258 401.423 297.886C402.09 297.508 402.991 297.318 404.128 297.318C405.598 297.318 406.893 297.686 408.014 298.42C409.135 299.155 410.01 300.197 410.639 301.545C411.268 302.894 411.582 304.485 411.582 306.318C411.582 308.167 411.268 309.769 410.639 311.125C410.01 312.473 409.139 313.519 408.026 314.261C406.912 314.996 405.628 315.364 404.173 315.364C403.052 315.364 402.154 315.178 401.48 314.807C400.806 314.428 400.287 314 399.923 313.523C399.56 313.038 399.279 312.636 399.082 312.318H398.855V321.545H396.173ZM398.81 306.273C398.81 307.591 399.003 308.754 399.389 309.761C399.776 310.761 400.34 311.545 401.082 312.114C401.825 312.674 402.734 312.955 403.81 312.955C404.931 312.955 405.866 312.659 406.616 312.068C407.374 311.47 407.942 310.667 408.321 309.659C408.707 308.644 408.901 307.515 408.901 306.273C408.901 305.045 408.711 303.939 408.332 302.955C407.961 301.962 407.397 301.178 406.639 300.602C405.889 300.019 404.946 299.727 403.81 299.727C402.719 299.727 401.802 300.004 401.06 300.557C400.317 301.102 399.757 301.867 399.378 302.852C398.999 303.83 398.81 304.97 398.81 306.273ZM418.355 291.727V315H415.673V291.727H418.355ZM430.358 315.364C428.782 315.364 427.4 314.989 426.21 314.239C425.028 313.489 424.104 312.439 423.438 311.091C422.778 309.742 422.449 308.167 422.449 306.364C422.449 304.545 422.778 302.958 423.438 301.602C424.104 300.246 425.028 299.193 426.21 298.443C427.4 297.693 428.782 297.318 430.358 297.318C431.934 297.318 433.313 297.693 434.494 298.443C435.684 299.193 436.608 300.246 437.267 301.602C437.934 302.958 438.267 304.545 438.267 306.364C438.267 308.167 437.934 309.742 437.267 311.091C436.608 312.439 435.684 313.489 434.494 314.239C433.313 314.989 431.934 315.364 430.358 315.364ZM430.358 312.955C431.555 312.955 432.54 312.648 433.312 312.034C434.085 311.42 434.657 310.614 435.028 309.614C435.4 308.614 435.585 307.53 435.585 306.364C435.585 305.197 435.4 304.11 435.028 303.102C434.657 302.095 434.085 301.28 433.312 300.659C432.54 300.038 431.555 299.727 430.358 299.727C429.161 299.727 428.176 300.038 427.403 300.659C426.631 301.28 426.059 302.095 425.688 303.102C425.316 304.11 425.131 305.197 425.131 306.364C425.131 307.53 425.316 308.614 425.688 309.614C426.059 310.614 426.631 311.42 427.403 312.034C428.176 312.648 429.161 312.955 430.358 312.955ZM447.497 315.409C446.391 315.409 445.387 315.201 444.486 314.784C443.584 314.36 442.868 313.75 442.338 312.955C441.808 312.152 441.543 311.182 441.543 310.045C441.543 309.045 441.74 308.235 442.134 307.614C442.527 306.985 443.054 306.492 443.713 306.136C444.372 305.78 445.099 305.515 445.895 305.341C446.698 305.159 447.505 305.015 448.315 304.909C449.376 304.773 450.236 304.67 450.895 304.602C451.562 304.527 452.046 304.402 452.349 304.227C452.66 304.053 452.815 303.75 452.815 303.318V303.227C452.815 302.106 452.509 301.235 451.895 300.614C451.289 299.992 450.368 299.682 449.134 299.682C447.853 299.682 446.849 299.962 446.122 300.523C445.395 301.083 444.884 301.682 444.588 302.318L442.043 301.409C442.497 300.348 443.103 299.523 443.861 298.932C444.626 298.333 445.459 297.917 446.361 297.682C447.27 297.439 448.164 297.318 449.043 297.318C449.603 297.318 450.247 297.386 450.974 297.523C451.709 297.652 452.418 297.92 453.099 298.33C453.789 298.739 454.361 299.356 454.815 300.182C455.27 301.008 455.497 302.114 455.497 303.5V315H452.815V312.636H452.679C452.497 313.015 452.194 313.42 451.77 313.852C451.346 314.284 450.781 314.652 450.077 314.955C449.372 315.258 448.512 315.409 447.497 315.409ZM447.906 313C448.967 313 449.861 312.792 450.588 312.375C451.323 311.958 451.876 311.42 452.247 310.761C452.626 310.102 452.815 309.409 452.815 308.682V306.227C452.702 306.364 452.452 306.489 452.065 306.602C451.687 306.708 451.247 306.803 450.747 306.886C450.255 306.962 449.774 307.03 449.304 307.091C448.842 307.144 448.467 307.189 448.179 307.227C447.482 307.318 446.83 307.466 446.224 307.67C445.626 307.867 445.141 308.167 444.77 308.568C444.406 308.962 444.224 309.5 444.224 310.182C444.224 311.114 444.569 311.818 445.259 312.295C445.955 312.765 446.838 313 447.906 313ZM466.983 315.364C465.528 315.364 464.244 314.996 463.131 314.261C462.017 313.519 461.146 312.473 460.517 311.125C459.888 309.769 459.574 308.167 459.574 306.318C459.574 304.485 459.888 302.894 460.517 301.545C461.146 300.197 462.021 299.155 463.142 298.42C464.263 297.686 465.559 297.318 467.028 297.318C468.165 297.318 469.063 297.508 469.722 297.886C470.388 298.258 470.896 298.682 471.244 299.159C471.6 299.629 471.877 300.015 472.074 300.318H472.301V291.727H474.983V315H472.392V312.318H472.074C471.877 312.636 471.597 313.038 471.233 313.523C470.869 314 470.35 314.428 469.676 314.807C469.002 315.178 468.104 315.364 466.983 315.364ZM467.347 312.955C468.422 312.955 469.331 312.674 470.074 312.114C470.816 311.545 471.381 310.761 471.767 309.761C472.153 308.754 472.347 307.591 472.347 306.273C472.347 304.97 472.157 303.83 471.778 302.852C471.4 301.867 470.839 301.102 470.097 300.557C469.354 300.004 468.438 299.727 467.347 299.727C466.21 299.727 465.263 300.019 464.506 300.602C463.756 301.178 463.191 301.962 462.812 302.955C462.441 303.939 462.256 305.045 462.256 306.273C462.256 307.515 462.445 308.644 462.824 309.659C463.21 310.667 463.778 311.47 464.528 312.068C465.286 312.659 466.225 312.955 467.347 312.955Z"
          fill="#EDECEF"
          id="svgButtonText"
        />
      </g>
      <g filter="url(#filter1_d_41_728)" id="svgMouse">
        <path
          d="M824.735 562.88C824.118 562.211 823 562.648 823 563.558V594.253C823 595.159 824.108 595.598 824.729 594.937L830.515 588.778L835.76 602.012C835.973 602.548 836.592 602.794 837.114 602.549L840.534 600.948C841.006 600.726 841.229 600.179 841.046 599.69L835.943 586.111H843.889C844.761 586.111 845.215 585.074 844.624 584.433L824.735 562.88Z"
          fill="#D9D9D9"
        />
        <path
          d="M824.735 562.88C824.118 562.211 823 562.648 823 563.558V594.253C823 595.159 824.108 595.598 824.729 594.937L830.515 588.778L835.76 602.012C835.973 602.548 836.592 602.794 837.114 602.549L840.534 600.948C841.006 600.726 841.229 600.179 841.046 599.69L835.943 586.111H843.889C844.761 586.111 845.215 585.074 844.624 584.433L824.735 562.88Z"
          stroke="black"
          strokeWidth="2"
        />
      </g>
      <defs>
        <filter
          id="filter0_d_41_728"
          x="299"
          y="271"
          width="250"
          height="73"
          filterUnits="userSpaceOnUse"
          colorInterpolationFilters="sRGB"
        >
          <feFlood floodOpacity="0" result="BackgroundImageFix" />
          <feColorMatrix
            in="SourceAlpha"
            type="matrix"
            values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"
            result="hardAlpha"
          />
          <feOffset dy="4" />
          <feGaussianBlur stdDeviation="2" />
          <feComposite in2="hardAlpha" operator="out" />
          <feColorMatrix
            type="matrix"
            values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0.25 0"
          />
          <feBlend
            mode="normal"
            in2="BackgroundImageFix"
            result="effect1_dropShadow_41_728"
          />
          <feBlend
            mode="normal"
            in="SourceGraphic"
            in2="effect1_dropShadow_41_728"
            result="shape"
          />
        </filter>
        <filter
          id="filter1_d_41_728"
          x="818"
          y="561.554"
          width="31.8928"
          height="50.0901"
          filterUnits="userSpaceOnUse"
          colorInterpolationFilters="sRGB"
        >
          <feFlood floodOpacity="0" result="BackgroundImageFix" />
          <feColorMatrix
            in="SourceAlpha"
            type="matrix"
            values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"
            result="hardAlpha"
          />
          <feOffset dy="4" />
          <feGaussianBlur stdDeviation="2" />
          <feComposite in2="hardAlpha" operator="out" />
          <feColorMatrix
            type="matrix"
            values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0.25 0"
          />
          <feBlend
            mode="normal"
            in2="BackgroundImageFix"
            result="effect1_dropShadow_41_728"
          />
          <feBlend
            mode="normal"
            in="SourceGraphic"
            in2="effect1_dropShadow_41_728"
            result="shape"
          />
        </filter>
        <clipPath id="clip0_41_728">
          <rect width="848" height="608" rx="16" fill="white" />
        </clipPath>
      </defs>
    </svg>
  );
}

export default LandingHero;
