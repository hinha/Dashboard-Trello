export const ClusterColumn = [
  {
    id: 1,
    name: "No",
    selector: (row) => row.no,
    sortable: true,
    reorder: true,
  },
  {
    id: 2,
    name: "Month",
    selector: (row) => row.time,
    sortable: true,
    reorder: true,
  },
  {
    id: 3,
    name: "Complete",
    selector: (row) => row.complete,
    sortable: true,
    reorder: true,
  },
  {
    id: 4,
    name: "Active",
    selector: (row) => row.active,
    sortable: true,
    right: true,
    reorder: true,
  },
  {
    name: "Year",
    selector: (row) => row.years,
    sortable: true,
    right: true,
    reorder: true,
  },
  {
    id: 6,
    name: "Cluster 1",
    selector: (row) => row.cluster_1,
    sortable: true,
    right: true,
    reorder: true,
  },
  {
    id: 7,
    name: "Cluster 2",
    selector: (row) => row.cluster_2,
    sortable: true,
    right: true,
    reorder: true,
  },
  {
    id: 8,
    name: "Cluster 3",
    selector: (row) => row.cluster_3,
    sortable: true,
    right: true,
    reorder: true,
  },
  {
    id: 9,
    name: "Cluster 4",
    selector: (row) => row.cluster_4,
    sortable: true,
    right: true,
    reorder: true,
  },
  {
    id: 10,
    name: "Distance",
    selector: (row) => row.distance,
    sortable: true,
    right: true,
    reorder: true,
  },
  {
    id: 11,
    name: "Structure",
    selector: (row) => row.structure,
    sortable: true,
    right: true,
    reorder: true,
  },
];

export const ClusterCoefficient = (averageData = []) => {
  let label = [];
  let data = [];
  averageData.forEach((item) => {
    label.push(`Cluster ${item.i}`);

    let unf_value = 0;
    if (item.v !== undefined) {
      unf_value = item.v;
    }

    data.push(unf_value);
  });
  return {
    grid: {
      left: "3%",
      right: "4%",
      top: "3%",
      bottom: "6%",
    },
    tooltip: {
      show: true,
    },
    xAxis: {
      type: "category",
      data: label,
    },
    yAxis: {
      type: "value",
    },
    series: [
      {
        data: data,
        type: "line",
      },
    ],
  };
};
