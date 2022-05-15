package seeds

import (
	"log"

	"github.com/joshua-ather/sv_be/app/models"
)

func (s Seed) PostSeed() {
	var err error
	for i, _ := range posts {
		err = s.db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed post table: %v", err)
		}
	}
}

var posts = []models.Post{
	models.Post{
		Title:    "Bitcoin",
		Content:  "Bitcoin (₿) is a decentralized digital currency that can be transferred on the peer-to-peer bitcoin network. Bitcoin transactions are verified by network nodes through cryptography and recorded in a public distributed ledger called a blockchain. The cryptocurrency was invented in 2008 by an unknown person or group of people using the name Satoshi Nakamoto. The currency began use in 2009[11] when its implementation was released as open-source software.: ch. 1 \n\nBitcoins are created as a reward for a process known as mining. They can be exchanged for other currencies, products, and services. Bitcoin has been criticized for its use in illegal transactions, the large amount of electricity (and thus carbon footprint) used by mining, price volatility, and thefts from exchanges. Some investors and economists have characterized it as a speculative bubble at various times. Others have used it as an investment, although several regulatory agencies have issued investor alerts about bitcoin.[12][13][14] Bitcoin, has been described as an economic bubble by at least eight Nobel Memorial Prize in Economic Sciences laureates, including Robert Shiller,[15] Joseph Stiglitz, and Richard Thaler. Journalists, economists, investors, and the central bank of Estonia have voiced concerns that bitcoin is a Ponzi scheme.\n\nThe word bitcoin was defined in a white paper published on 31 October 2008. It is a compound of the words bit and coin. No uniform convention for bitcoin capitalization exists; some sources use Bitcoin, capitalized, to refer to the technology and network and bitcoin, lowercase, for the unit of account. The Wall Street Journal, The Chronicle of Higher Education, and the Oxford English Dictionary advocate the use of lowercase bitcoin in all cases.\n\nA few local and national governments are officially using Bitcoin in some capacity, with two countries, El Salvador and the Central African Republic, adopting it as a legal tender.",
		Category: "Cryptocurrency",
		Status:   "Publish",
	},
	models.Post{
		Title:    "Semantic Web (Web 3.0)",
		Content:  "The Semantic Web, sometimes known as Web 3.0, is an extension of the World Wide Web through standards[1] set by the World Wide Web Consortium (W3C). The goal of the Semantic Web is to make Internet data machine-readable.\n\nTo enable the encoding of semantics with the data, technologies such as Resource Description Framework (RDF) and Web Ontology Language (OWL)[3] are used. These technologies are used to formally represent metadata. For example, ontology can describe concepts, relationships between entities, and categories of things. These embedded semantics offer significant advantages such as reasoning over data and operating with heterogeneous data sources.\n\nThese standards promote common data formats and exchange protocols on the Web, fundamentally the RDF. According to the W3C, \"The Semantic Web provides a common framework that allows data to be shared and reused across application, enterprise, and community boundaries.\" The Semantic Web is therefore regarded as an integrator across different content and information applications and systems.\n\nThe term was coined by Tim Berners-Lee for a web of data (or data web)[6] that can be processed by machines[7]—that is, one in which much of the meaning is machine-readable. While its critics have questioned its feasibility, proponents argue that applications in library and information science, industry, biology and human sciences research have already proven the validity of the original concept.",
		Category: "Web",
		Status:   "Draft",
	},
	models.Post{
		Title:    "5G Technology",
		Content:  "In telecommunications, 5G is the fifth-generation technology standard for broadband cellular networks, which cellular phone companies began deploying worldwide in 2019, and is the planned successor to the 4G networks which provide connectivity to most current cellphones. 5G networks are predicted to have more than 1.7 billion subscribers worldwide by 2025, according to the GSM Association.\n\nLike its predecessors, 5G networks are cellular networks, in which the service area is divided into small geographical areas called cells. All 5G wireless devices in a cell are connected to the Internet and telephone network by radio waves through a local antenna in the cell. The new networks have greater bandwidth than their predecessors, giving higher download speeds, eventually up to 10 gigabits per second (Gbit/s). In addition to 5G being faster than existing networks, 5G has higher bandwidth and can thus connect more different devices, improving the quality of Internet services in crowded areas. Due to the increased bandwidth, it is expected the networks will increasingly be used as general internet service providers (ISPs) for laptops and desktop computers, competing with existing ISPs such as cable internet, and also will make possible new applications in internet-of-things (IoT) and machine-to-machine areas. Cellphones with 4G capability alone are not able to use the new networks, which require 5G-enabled wireless devices.",
		Category: "Network",
		Status:   "Thrash",
	},
}
