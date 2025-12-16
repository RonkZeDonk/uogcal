import MainLayout from "@/layouts/main";
import { Button, Group, Stack, Text, Title } from "@mantine/core";
import { SelectionBox } from "@/components/SelectionBox";
import { useState } from "react";
import { useNavigate } from "react-router-dom";

// I fetched these values with this command:
// $ curl -s https://colleague-ss.uoguelph.ca/Student/Courses/GetCatalogAdvancedSearch | jq -c '[.Subjects[].Code]'
// TODO dynamically fetch this on the server side every couple days/months
const subjects = ["ACCT","AGR","DAGR","ANSC","ANTH","ARAB","ARTH","AVC","ASCI","AHSS","BIOC","BINF","BIOL","BIOM","BIOP","BIOT","BLCK","BOT","BUS","BADM","CDE","CHEM","CHIN","CLAS","CLIN","CSS","CIS","CONS","COOP","CREA","CRWR","CJPP","CCJP","IMPR","CROP","CTS","DATA","ECS","ECON","ENGG","ENGL","EDRD","ENVM","DENM","ENVS","EQN","DEQN","EURO","FCSS","FRHD","FRAN","FIN","FINA","FSQA","FOOD","FARE","FREN","GEOG","GERM","GREK","HISP","HIST","HORT","DHRT","HTM","HHNS","HK","HROB","HUMN","IES","INDG","IBIO","IAEF","IPS","ISS","UNIV","IDEV","ITAL","JLS","JUST","KIN","LARC","LAT","LACS","LEAD","LING","LTS","MGMT","MCS","MATH","MDST","MICR","MCB","MBG","MUSC","NANO","NEUR","NUTR","ONEH","OAGR","PABI","PATH","CPHH","PHIL","PHYS","PLNT","PBIO","POLS","POPM","PORT","PSYC","REAL","ROY","RPD","RST","SCMA","XSEN","SXGN","SOPR","SOC","SOAN","SPAN","SPMT","STAT","SART","THST","TRMH","TOX","DTM","VETM","CVOA","DVT","WMST","ZOO"];

function ImportPage() {
  // const navigate = useNavigate();
  // fetch("upload/courses", { method: "post", body: fd }).then(() => navigate("/account"));

  const navigate = useNavigate();
  const [order, setOrder] = useState<
    { sub?: string; code?: string; sect?: string }[]
  >([]);

  return (
    <MainLayout>
      <div className="px-8 py-12 sm:px-16">
        <Title>Import your courses</Title>

        {/* TODO a list of inputted data */}

        {order.map((e, idx) => (
          <Group gap={0} key={idx}>
            <SelectionBox
              value={e.sub}
              label="Subject"
              placeholder="MATH"
              options={subjects}
              onSelect={(v) =>
                setOrder((x) => [...x.slice(0, idx), { sub: v || "" }])
              }
              side="left"
            />
            <SelectionBox
              value={e.code}
              label="Course code"
              placeholder="1200"
              onSelect={(v) =>
                setOrder((x) => [...x.slice(0, idx), { ...x[idx], code: v || "" }])
              }
            />
            <SelectionBox
              value={e.sect}
              label="Section"
              placeholder="0101"
              onSelect={(v) =>
                setOrder((x) => [...x.slice(0, idx), { ...x[idx], sect: v || "" }])
              }
              side="right"
            />
          </Group>
        ))}

        <Button onClick={() => setOrder((old) => [...old, {}])}>
          + Add course
        </Button>
        <Button
          onClick={() => {
            const data = {
              term: "F25",
              courses: order.map((v) => ({
                title: `${v.sub}*${v.code}*${v.sect}`
              })),
            };
            const fd = new FormData();
            fd.append(
              "courses",
              new Blob([JSON.stringify(data)], { type: "application/json" }),
              "courses.json"
            );
            fetch("upload/courses", { method: "post", body: fd }).then(() =>
              navigate("/account")
            );
          }}
        >
          Submit
        </Button>
      </div>
    </MainLayout>
  );
}
  // fetch("upload/courses", { method: "post", body: fd }).then(() => navigate("/account"));

export default ImportPage;
