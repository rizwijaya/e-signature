// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity ^0.8.0;
pragma abicoder v2;

contract Smartsign {
    address public sistem;
    address public userAccess;
    uint256 profileCount;
    uint256 signatureCount;
    uint256 documentCount;
    uint256 mySignatureCount;

    constructor() {
        sistem = msg.sender;
        userAccess = msg.sender;
        profileCount = 1;
        signatureCount = 1;
        documentCount = 1;
        mySignatureCount = 1;
    }

    //------Model data------//
    //Modal Signature Saya
    struct MySignature {
        address signers;
        string signature_data;
        string signature_nodata;
        string latin_data;
        string latin_nodata;
        string date_updated;
    }

    //Model Document Data
    struct Document {
        //Document selama proses ttd dari awal sampai selesai
        uint256 document_id;
        bytes32 hash_original; //identifier hash original file untuk get data history ttd dan urutkan paling baru
        address signers;
        uint signing; //Yang telah ttd
        string namefile;
        string hash_file;
        string hash_ipfs;
        string metadata;
        uint state; //1 Process Signed, 2 Signed
        uint totalsigned;
        uint createdtime;
        uint completedtime;
        bool exist;
        mapping(address => Allsigners) allsigners;
    }
    //Model Document Signers
    struct Allsigners {
        address signer_address;
        bool state; //True is signed and False is not signed
        string date; //Tanggal TTD
        bool exist;
    }
    //------End Model data------//

    //------Permission Access------//
    // Permission hanya admin/sistem yang dapat mengakses
    modifier onlySistem() {
        require(msg.sender == sistem);
        _;
    }
    //Permission User yang dapat mengakses
    modifier permissionUser() {
        require(msg.sender == userAccess);
        _;
    }
    //Multipermissions
    mapping(address => bool) members;

    function addPermission(address _member) public {
        members[_member] = true;
    }

    function permission(address _member) internal view returns (bool) {
        if (members[_member]) {
            return true;
        }
        return false;
    }

    //Document Permission for Signing
    function docPermission(bytes32 _hash) internal view returns (bool) {
        if (documents[_hash].exist == false) {
            return false;
        }
        if (documents[_hash].signing == 0 || documents[_hash].state == 2) {
            //Tidak ada dokumen atau Telah melakukan ttd semua
            return false;
        }
        //Permission Address Signed or Not
        if (documents[_hash].allsigners[msg.sender].exist == true) {
            //Terdapat Address
            if (documents[_hash].allsigners[msg.sender].state == false) {
                //Jika address belum ttd
                return true;
            }
        }
        return false;
    }

    //------End Permission Access------//

    //------MySignatures Function------//
    address[] public mySignature;
    mapping(address => MySignature) public mySign;
    //Add My-Signatures
    function addmysignatures(string memory _signature_data, string memory _signature_nodata, string memory _latin_data, string memory _latin_nodata, string memory _date_updated) public {
        MySignature memory newMySignature = 
            MySignature({
                signers: msg.sender,
                signature_data: _signature_data,
                signature_nodata: _signature_nodata,
                latin_data: _latin_data,
                latin_nodata: _latin_nodata,
                date_updated: _date_updated
            });
            mySign[msg.sender] = newMySignature;
            mySignature.push(msg.sender);
            mySignatureCount += 1;
    }
    //------End MySignatures Function------//

    //------Document Function------//
    // mapping(bytes32 => Document) documents;
    bytes32[] public documentlists;
    mapping(bytes32 => Document) documents;

    // mapping(bytes32 => Document) documents;
    // bytes32[] documentlist;
    //------Signing Document------//
    function signDoc(
        bytes32 _hash_original,
        string memory _namefile,
        string memory _hash_file,
        string memory _metadata,
        string memory _hash_ipfs,
        uint256 _totalsigned,
        uint256 signing,
        uint256 _createdtime
    ) public {
        require(
            docPermission(_hash_original),
            "You are not allowed to sign this document"
        );
        if (documents[_hash_original].signing - 1 == 0) {
            //User terakhir yang melakukan ttd
            documents[_hash_original].state = 2; //State Signed
        } else {
            documents[_hash_original].state = 1; //state Process Sign
        }

        documents[_hash_original].document_id = documentCount;
        documents[_hash_original].signers = msg.sender;
        documents[_hash_original].namefile = _namefile;
        documents[_hash_original].hash_original = _hash_original;
        documents[_hash_original].hash_file = _hash_file;
        documents[_hash_original].metadata = _metadata;
        documents[_hash_original].hash_ipfs = _hash_ipfs;
        documents[_hash_original].signing = signing - 1;
        documents[_hash_original].totalsigned = _totalsigned;
        documents[_hash_original].createdtime = _createdtime;
        documents[_hash_original].completedtime = block.timestamp;
        documents[_hash_original].exist = true;
        documentlists.push(_hash_original);
        documentCount += 1;
    }

    //Add signers in ttd process
    function addsigners(
        bytes32 _hash_original,
        address _signers,
        string memory _date
    ) public {
        require(
            docPermission(_hash_original),
            "You are not allowed to sign this document"
        );
        documents[_hash_original].allsigners[_signers] = Allsigners({
            signer_address: _signers,
            state: true,
            date: _date,
            exist: true
        });
    }

    //------End Signing Document------//
    //------Get Document Latest------//

    //------End Get Document Latest------//
    //------Verify Document Function------//
    function verify_document(bytes32 _hash) public view returns (bool) {
        if (keccak256(bytes(documents[_hash].hash_file)) != keccak256(bytes(""))) {
            //Document Exist
            return true;
        }
        return false;
    }
    // ------End Document Function------//
}
