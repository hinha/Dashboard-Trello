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
    name: "Date",
    selector: (row) => row.created_at,
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
    name: "Comment",
    selector: (row) => row.comment,
    sortable: true,
    right: true,
    reorder: true,
  },
  {
    id: 5,
    name: "Total",
    selector: (row) => row.total,
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
    name: "Cluster 5",
    selector: (row) => row.cluster_5,
    sortable: true,
    right: true,
    reorder: true,
  },
  {
    id: 11,
    name: "Cluster 6",
    selector: (row) => row.cluster_6,
    sortable: true,
    right: true,
    reorder: true,
  },
  {
    id: 12,
    name: "Cluster 7",
    selector: (row) => row.cluster_7,
    sortable: true,
    right: true,
    reorder: true,
  },
  {
    id: 13,
    name: "Cluster 8",
    selector: (row) => row.cluster_8,
    sortable: true,
    right: true,
    reorder: true,
  },
  {
    id: 14,
    name: "Cluster 9",
    selector: (row) => row.cluster_9,
    sortable: true,
    right: true,
    reorder: true,
  },
  {
    id: 15,
    name: "Cluster 10",
    selector: (row) => row.cluster_10,
    sortable: true,
    right: true,
    reorder: true,
  },
];

export const ClusterCoefficient = (averageData = []) => {
  console.log(averageData);
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
